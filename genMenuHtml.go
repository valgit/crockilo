package main

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
	"os"
)

type Recette struct {
	Name    string
	Portion int
	Img     string
	Link    string
}
type MenuData struct {
	Name string
	Plat []Recette
}

type DayData struct {
	Meal string
	Menu []MenuData
}

type WeekData struct {
	Title string
	Day   []DayData
}

func GenHTML(weekData WeekData, tmpl string, outputPath string) {
	// Files are provided as a slice of strings.

	parsedTemplate, _ := template.ParseFiles(tmpl)

	// create file
	var processed bytes.Buffer

	err := parsedTemplate.Execute(&processed, weekData)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}

	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()
}
