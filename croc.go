package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func setHeaders(req *http.Request) {
	req.Header.Set("authority", "api.croq-kilos.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("content-type", "null")
	req.Header.Set("dnt", "1")
	req.Header.Set("origin", "https://app.croq-kilos.com")
	req.Header.Set("referer", "https://app.croq-kilos.com/")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
}

func setJWTHeaders(req *http.Request, jwt string) {
	token := fmt.Sprintf("Bearer %s", jwt)
	req.Header.Set("authorization", token)

}

func CheckAuth(user string, pass string) (Jwtlogin, error) {
	client := &http.Client{}

	loginPost := map[string]string{
		"username": user,
		"password": pass,
	}
	data, err := json.Marshal(loginPost)
	if err != nil {
		return Jwtlogin{}, err
	}

	req, err := http.NewRequest("POST", "https://api.croq-kilos.com/login_check", bytes.NewBuffer(data))
	if err != nil {
		return Jwtlogin{}, err
	}

	req.Header.Set("authority", "api.croq-kilos.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "fr-FR,fr;q=0.6")
	req.Header.Set("content-type", "application/json;")
	req.Header.Set("origin", "https://app.croq-kilos.com")
	req.Header.Set("referer", "https://app.croq-kilos.com/")
	req.Header.Set("sec-ch-ua", `"Brave";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return Jwtlogin{}, err
	}
	defer resp.Body.Close()

	//JSON unmarshal
	var authInfo Jwtlogin

	//use alternativ unmarshal
	err = json.NewDecoder(resp.Body).Decode(&authInfo)
	if err != nil {
		return Jwtlogin{}, err
	}

	return authInfo, nil
}

/* param : time since epoch of require day
 * 1674514800 => 23/01/ @ 23h00 GMT
 */
func GetMenuForDay(day int64, JWT string) DaylyMenu {
	client := &http.Client{}

	var baseUrl = fmt.Sprintf("https://api.croq-kilos.com/menus?day=%d", day)

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	setHeaders(req)

	setJWTHeaders(req, JWT)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		log.Fatal(resp.Status)
	}

	//fmt.Printf("body:\n%sın", resp.Body)

	//JSON unmarshal
	var infomenu DaylyMenu

	//use alternativ unmarshal
	err = json.NewDecoder(resp.Body).Decode(&infomenu)
	if err != nil {
		log.Fatal(err)
	}

	//TODO: log for debug
	//fmt.Printf("%+v\n", infomenu)
	return infomenu
}

/* param : id of recipe
 */
func GetRecipe(recipeid string, JWT string) MealRecipe {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.croq-kilos.com"+recipeid, nil)
	if err != nil {
		log.Fatal(err)
	}

	setHeaders(req)

	setJWTHeaders(req, JWT)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	/*
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", bodyText)
	*/
	var recipe MealRecipe

	//use alternativ unmarshal
	err = json.NewDecoder(resp.Body).Decode(&recipe)
	if err != nil {
		log.Fatal(err)
	}

	return recipe
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
}

func generateWeekMenu(date time.Time) string {
	cwd, err := GetWorkingDir() // os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}

	fmt.Println("Current working directory:", cwd)
	defaultConf := cwd + "/config.yml"
	configFile := flag.String("config_file", defaultConf, "configuration file path")
	var user string
	var pass string
	var basedir string

	flag.StringVar(&user, "u", "", "user name")
	flag.StringVar(&pass, "p", "", "password")
	flag.StringVar(&basedir, "d", "", "base directory")
	flag.Usage = usage
	flag.Parse()

	cfg, err := LoadConfig(*configFile)
	if err != nil {
		log.Printf("error loading config file: %v", err)
		return ""
	}

	// Override config values with command-line flags
	if user != "" {
		cfg.Appconf.User = user
	}
	if pass != "" {
		cfg.Appconf.Password = pass
	}
	if basedir != "" {
		cfg.Appconf.Basedir = basedir
	}

	// Check required parameters
	if cfg.Appconf.User == "" || cfg.Appconf.Password == "" {
		flag.Usage()
		os.Exit(1)
	}

	/* get a one time token*/
	token, err := CheckAuth(cfg.Appconf.User, cfg.Appconf.Password)
	if err != nil {
		log.Fatal("login error")
		return ""
	}
	fmt.Printf("JWT: %s\n", token.Token)
	/**/

	fmt.Printf("basedir is : %s\n", cfg.Appconf.Basedir)
	//pathmenu := GetDocuments() + "/crocmenu"
	pathmenu := cfg.Appconf.Basedir + "/crocmenu"
	err = Checkdir(pathmenu)
	if err != nil {
		log.Fatal("can't create menu dir")
		return ""
	} else {
		fmt.Printf("will store menu in %s\n", pathmenu)
	}
	// recipe
	pathrecipe := pathmenu + "/recipes"
	err = Checkdir(pathrecipe)
	if err != nil {
		log.Fatal("can't create recipe dir")
		return ""
	} else {
		fmt.Printf("will store recipe in %s\n", pathrecipe)
	}

	fmt.Printf("will get menu for %s\n", date)
	// today
	humanReadable := time.Now()
	// Date de départ
	//date := time.Date(2023, time.April, 2, 0, 0, 0, 0, time.UTC)

	// Trouver le premier dimanche précédent la date de départ
	startOfWeek := humanReadable.AddDate(0, 0, -int(humanReadable.Weekday()))
	fmt.Printf("start : %s\n", startOfWeek)

	year, week := humanReadable.ISOWeek()
	//
	_, w := date.ISOWeek()
	fmt.Printf("%d %d\n", week, w)
	outputPath := fmt.Sprintf("%s/menu_%d_%d.html", pathmenu, year, week)

	weekData := WeekData{
		Title: fmt.Sprintf("Menu de la Semaine %d", week),
		Day:   make([]DayData, 7),
	}

	/* compile recipe tmpl */
	reciptmpl := pathmenu + "/recipe.html"
	if !IsFileExist(reciptmpl) {
		log.Fatal("template not found\n")
	}
	parsedTemplate, _ := template.ParseFiles(reciptmpl)

	for i := 0; i < 7; i++ {
		// Calculez la date du jour en utilisant la date courante et le numéro de jour
		//humanReadable := humanReadable.AddDate(0, 0, i-int(humanReadable.Weekday()))
		humanReadable := startOfWeek.AddDate(0, 0, i)

		day := humanReadableToEpoch24(humanReadable)

		fmt.Printf("time : %d %s\n", day, humanReadable)
		//fmt.Fprintf(f, "<h1>%s</h1>\n", humanReadable.Weekday())
		// localized
		r := strings.NewReplacer(
			"Monday", "Lundi",
			"Tuesday", "Mardi",
			"Wednesday", "Mercredi",
			"Thursday", "Jeudi",
			"Friday", "Vendredi",
			"Saturday", "Samedi",
			"Sunday", "Dimanche",
		)

		//weekday := fmt.Sprintf("%s", humanReadable.Weekday())
		weekday := humanReadable.Weekday().String()
		weekData.Day[i].Meal = r.Replace(weekday) // fmt.Sprintf("%s", weekday)
		weekData.Day[i].Menu = make([]MenuData, 4)

		menus := GetMenuForDay(day, token.Token)

		fmt.Printf("type : %s %d\n", menus.Type, menus.HydraTotalItems)

		for _, menu := range menus.HydraMember {
			fmt.Printf("type: %s \n", menu.Type)
			for m, meals := range menu.MenuMeals {
				//fmt.Fprintf(f, "<h1>%s</h1>\n", meals.Meal.Title)
				pieces := strings.Split(meals.Meal.Title, " - ")
				if len(pieces) > 2 {
					weekData.Day[i].Menu[m].Name = fmt.Sprintf("%s (%s)", pieces[1], pieces[2])
					fmt.Printf("menu: %s - %s\n", pieces[1], pieces[2])
				} else {
					weekData.Day[i].Menu[m].Name = fmt.Sprintf("%s", pieces[1])
					fmt.Printf("menu: %s \n", pieces[1])
				}
				//TODO: better
				weekData.Day[i].Menu[m].Plat = make([]Recette, 5)

				for p, recipes := range meals.Meal.MealRecipes {
					//fmt.Fprintf(f, "<h2>%s</h2>\n", recipes.Recipe.Title)
					fmt.Printf("%s\n", recipes.Recipe.Title)

					weekData.Day[i].Menu[m].Plat[p].Name = recipes.Recipe.Title

					/*
						fmt.Printf("portion: %d\nétapes :\n %s\n", recipes.Recipe.Portion,
							recipes.Recipe.Steps)
					*/
					weekData.Day[i].Menu[m].Plat[p].Portion = recipes.Recipe.Portion
					if recipes.Recipe.Steps != "" {

						recipe := GetRecipe(recipes.Recipe.ID, token.Token)
						//TODO: save file

						weekData.Day[i].Menu[m].Plat[p].Img =
							fmt.Sprintf("https://api.croq-kilos.com/%s",
								recipe.Picture.WebPath)

						recipeData := RecipeData{
							Title:       recipe.Title,
							Img:         fmt.Sprintf("https://api.croq-kilos.com/%s", recipe.Picture.WebPath),
							Portion:     recipe.Portion,
							CookingTime: recipe.CookingTime,
							Kcal:        recipe.Kcal,
							Steps:       template.HTML(recipe.Steps),
						}

						//recipeData.Ingredients = make([]Ingredient, 5)

						for i, ingredient := range recipe.RecipeIngredients {
							recipeData.Ingredients[i].Name = ingredient.Ingredient.Title
							recipeData.Ingredients[i].Quantity = int64(ingredient.Quantity)
							recipeData.Ingredients[i].Unit = GetUnit(ingredient.Ingredient.Unit)
						}

						fname := fmt.Sprintf("%s/%d.html", pathrecipe, recipe.ID0)
						GenRecipeHTML(recipeData, parsedTemplate, fname)
						fname = fmt.Sprintf("recipes/%d.html", recipe.ID0)
						weekData.Day[i].Menu[m].Plat[p].Link = fname
					} else {
						weekData.Day[i].Menu[m].Plat[p].Link = ""
					}
					//fmt.Fprintf(f, "<img src=\"https://api.croq-kilos.com/%s\" width=200></img>", recipe.Picture.WebPath)
					/*fmt.Fprintf(f, "<p>Portion: %d\ncuisson: %d, Kcal: %d\n",
						recipe.Portion, recipe.CookingTime, recipe.Kcal)
					if len(recipe.Steps) != 0 {
						fmt.Fprintf(f, "<h3>étapes</h3>\n%s\n", recipe.Steps)
					}*/

				}
				fmt.Println("----")
			}

		}
	}
	// output HTML data
	weektmpl := pathmenu + "/layout.html"
	if !IsFileExist(weektmpl) {
		log.Fatal("template not found\n")
	}
	GenHTML(weekData, weektmpl, outputPath)

	// reindex recette
	/*
		outputPath = fmt.Sprintf("%s/recettes.html", pathmenu)
		IndexMenu(pathrecipe, outputPath)
	*/
	return outputPath
}
