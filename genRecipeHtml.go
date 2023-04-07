package main

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
	"os"
)

type IngredientData struct {
	Name     string
	Quantity int64
	Unit     string
}

type RecipeData struct {
	Title       string
	Img         string
	Portion     int
	CookingTime int
	Kcal        int
	Steps       template.HTML
	Ingredients [20]IngredientData
}

/*
 * generate a recipe file
 */
func GenRecipeHTML(recipeData RecipeData, parsedTemplate *template.Template, outputPath string) {
	// Files are provided as a slice of strings.
	/*
		parsedTemplate, _ := template.ParseFiles(tmpl)
	*/
	// create file
	var processed bytes.Buffer

	err := parsedTemplate.Execute(&processed, recipeData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal("open error")
	}
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()
}
