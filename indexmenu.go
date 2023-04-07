package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type IndexData struct {
	FileNname []string
}

func genIndexHTML(indexData IndexData, tmpl string, outputPath string) {
	// Files are provided as a slice of strings.

	parsedTemplate, _ := template.ParseFiles(tmpl)

	// create file
	var processed bytes.Buffer

	err := parsedTemplate.Execute(&processed, indexData)
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

/*
 * create an HTML file with all the content
 */
func IndexMenu(dirPath string, filePath string) {
	// Read the content of the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	idxdata := IndexData{
		FileNname: make([]string, 10),
	}

	// Write the list of files to the output file
	for _, fileInfo := range files {
		idxdata.FileNname = append(idxdata.FileNname, fileInfo.Name())
	}

	genIndexHTML(idxdata, "files.html", filePath)
}
