package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type IndexData struct {
	FileName string
	Title    string
}

func genIndexHTML(indexData []IndexData, tmpl string, outputPath string) {
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
/*
func IndexMenuOld(dirPath string, filePath string) {
	// Read the content of the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	idxdata := IndexData{
		FileName: make([]string, 10),
	}

	// Write the list of files to the output file
	for _, fileInfo := range files {
		idxdata.FileName = append(idxdata.FileName, fileInfo.Name())
	}

	genIndexHTML(idxdata, "files.html", filePath)
}
*/

var titlePattern = regexp.MustCompile(`(?i)<title>(.*?)<\/title>`)

/**
 * extract title info
 **/
func extractTitleFromFile(filePath string) (string, error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	matches := titlePattern.FindSubmatch(fileContent)
	if len(matches) >= 2 {
		return string(matches[1]), nil
	}

	return "", fmt.Errorf("no title found in file: %s", filePath)
}

func indexDirectory(directoryPath string) (map[string]string, error) {
	index := make(map[string]string)

	err := filepath.Walk(directoryPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			title, err := extractTitleFromFile(filePath)
			if err != nil {
				log.Printf("Error extracting title from file: %s\n%s", filePath, err.Error())
				return nil // Skip the file and continue indexing
			}

			index[filePath] = title
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return index, nil
}

func IndexMenu(dirPath string, filePath string, tmpl string) error {
	//directoryPath := "/path/to/your/directory"
	if !IsFileExist(tmpl) {
		return errors.New("template file not found")
	}
	if !IsFileExist(dirPath) {
		return errors.New("index directory not found")
	}
	index, err := indexDirectory(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	// Print the indexed title
	var dataArray []IndexData

	for fileName, title := range index {
		data := IndexData{
			FileName: fileName,
			Title:    title,
		}
		dataArray = append(dataArray, data)
	}

	genIndexHTML(dataArray, tmpl, filePath)
	return nil
}
