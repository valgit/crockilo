package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

/*
 * some tooling
 */

/*
 * get user dir and documents
 */
func GetDocuments() string {
	documentsDir := os.Getenv("HOME") + "/Documents"
	//fmt.Println("Documents directory:", documentsDir)
	return documentsDir
}

/*
 * test if dir exist and create it if needed
 */
func Checkdir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Create directory if it doesn't exist
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

// Check if file exists
func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		//fmt.Printf("File %s does not exist\n", filename)
		return false
	} else {
		return true
		//fmt.Printf("File %s exists\n", filename)
	}
}

/*
 * save some url to file
 */
func SaveUrl(url string, fname string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File saved!")
}

func GetWorkingDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	programDir := filepath.Dir(exePath)
	return programDir, nil
}

/*
 * convert numer to real unit
 */
func GetUnit(unit string) string {
	switch unit {
	case "1":
		return "g"
	case "4":
		return "ml"
	default:
		return ""
	}
}

// "file:///Users/val/Documents/crockilo/crocmenu/menu_2023_18.html"
func getHTMLpath() (string, error) {
	cwd, err := GetWorkingDir() // os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return "", err
	}

	fmt.Println("Current working directory:", cwd)
	defaultConf := cwd + "/config.yml"
	cfg, err := LoadConfig(defaultConf)
	if err != nil {
		//log.Printf("error loading config file: %v", err)
		return "", err
	}

	fmt.Printf("basedir is : %s\n", cfg.Appconf.Basedir)
	//pathmenu := GetDocuments() + "/crocmenu"
	pathmenu := cfg.Appconf.Basedir + "/crocmenu"
	return pathmenu, nil
}
