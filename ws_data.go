package main

import "time"

/*
 * auth token
 */
type Jwtlogin struct {
	Token string `json:"token"`
	User  struct {
		ID         int      `json:"id"`
		Firstlogin int      `json:"firstlogin"`
		Roles      []string `json:"roles"`
	} `json:"user"`
	RefreshToken string `json:"refresh_token"`
}

/*
 * daily menu
 */
type DaylyMenu struct {
	Context         string        `json:"@context"`
	ID              string        `json:"@id"`
	Type            string        `json:"@type"`
	HydraMember     []HydraMember `json:"hydra:member"`
	HydraTotalItems int           `json:"hydra:totalItems"`
	HydraView       HydraView     `json:"hydra:view"`
	HydraSearch     HydraSearch   `json:"hydra:search"`
}
type Picture struct {
	ID           string `json:"@id"`
	Type         string `json:"@type"`
	WebPath      string `json:"webPath"`
	Medium       string `json:"medium"`
	Large        string `json:"large"`
	ArticleAsset string `json:"articleAsset"`
}
type Recipe struct {
	ID       string      `json:"@id"`
	Type     string      `json:"@type"`
	ID0      int         `json:"id"`
	Title    string      `json:"title"`
	Calories int         `json:"calories,omitempty"`
	Kcal     int         `json:"kcal"`
	Steps    string      `json:"steps"`
	Note     interface{} `json:"note"`
	Picture  Picture     `json:"picture"`
	VideoURL interface{} `json:"videoUrl"`
	Portion  int         `json:"portion"`
}
type MealRecipes struct {
	ID              string      `json:"@id"`
	Type            string      `json:"@type"`
	Recipe          Recipe      `json:"recipe"`
	Position        string      `json:"position"`
	User            interface{} `json:"user"`
	IsAlternative   interface{} `json:"isAlternative"`
	HasAlternatives interface{} `json:"hasAlternatives"`
}
type Meal struct {
	ID          string        `json:"@id"`
	Type        string        `json:"@type"`
	Title       string        `json:"title"`
	Type0       string        `json:"type"`
	MealRecipes []MealRecipes `json:"mealRecipes"`
}
type MenuMeals struct {
	ID              string        `json:"@id"`
	Type            string        `json:"@type"`
	Meal            Meal          `json:"meal"`
	MenuMealPersons []interface{} `json:"menuMealPersons"`
}
type HydraMember struct {
	ID          string      `json:"@id"`
	Type        string      `json:"@type"`
	ID0         string      `json:"id"`
	Title       interface{} `json:"title"`
	PublishedAt time.Time   `json:"publishedAt"`
	MenuMeals   []MenuMeals `json:"menuMeals"`
}
type HydraView struct {
	ID   string `json:"@id"`
	Type string `json:"@type"`
}
type HydraMapping struct {
	Type     string `json:"@type"`
	Variable string `json:"variable"`
	Property string `json:"property"`
	Required bool   `json:"required"`
}
type HydraSearch struct {
	Type                        string         `json:"@type"`
	HydraTemplate               string         `json:"hydra:template"`
	HydraVariableRepresentation string         `json:"hydra:variableRepresentation"`
	HydraMapping                []HydraMapping `json:"hydra:mapping"`
}

/*
 * meals recipes
 */
type MealRecipe struct {
	Context           string              `json:"@context"`
	ID                string              `json:"@id"`
	Type              string              `json:"@type"`
	ID0               int                 `json:"id"`
	Title             string              `json:"title"`
	Calories          int                 `json:"calories,omitempty"`
	Kcal              int                 `json:"kcal"`
	CreatedAt         time.Time           `json:"createdAt"`
	UpdatedAt         time.Time           `json:"updatedAt"`
	IsVegan           bool                `json:"isVegan"`
	AtTop             bool                `json:"atTop"`
	PreparationTime   int                 `json:"preparationTime"`
	CookingTime       int                 `json:"cookingTime"`
	Steps             string              `json:"steps"`
	Note              string              `json:"note"`
	Picture           Picture             `json:"picture"`
	Categories        []string            `json:"categories"`
	Seasons           []string            `json:"seasons"`
	RecipeIngredients []RecipeIngredients `json:"recipeIngredients"`
	IsStabilization   bool                `json:"isStabilization"`
	VideoURL          interface{}         `json:"videoUrl"`
	Rating            float64             `json:"rating"`
	RatingVolume      int                 `json:"ratingVolume"`
	Portion           int                 `json:"portion"`
	PublishedAtStart  time.Time           `json:"publishedAtStart"`
	PublishedAtEnd    interface{}         `json:"publishedAtEnd"`
	Status            int                 `json:"status"`
}

type Ingredient struct {
	ID    string `json:"@id"`
	Type  string `json:"@type"`
	ID0   int    `json:"id"`
	Title string `json:"title"`
	Unit  string `json:"unit"`
}
type RecipeIngredients struct {
	ID         string     `json:"@id"`
	Type       string     `json:"@type"`
	Ingredient Ingredient `json:"ingredient"`
	Quantity   float64    `json:"quantity"`
	Position   int        `json:"position"`
}
