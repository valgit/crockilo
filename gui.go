package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func openURL(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("crockilo")

	/*
		// Load the icon file
		iconResource, err := fyne.LoadResourceFromPath("croc.png")
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		// Set the icon for the application
		myApp.SetIcon(iconResource)
	*/

	// Set the date to today's date
	today := time.Now().Format("02/01/2006")
	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("Enter a date (dd/mm/yyyy)")
	dateEntry.SetText(today)

	checkbox := widget.NewCheck("start date", nil)

	outputLabel := widget.NewLabel("Output will appear here")

	startButton := widget.NewButton("Start", func() {
		dateStr := dateEntry.Text
		date, err := time.Parse("02/01/2006", dateStr)
		if err != nil {
			outputLabel.SetText("Error: Invalid date format")
			return
		}
		outputLabel.SetText(fmt.Sprintf("Starting process for date: %s", dateStr))

		// Call your function here with the date parameter
		// processDate(date)
		fmt.Printf("%s\n", date)
		getConfig()
		/* get a one time token*/
		token, err := getToken()
		if err != nil {
			outputLabel.SetText("login error")

		}
		menu := generateWeekMenu(date, token)
		err = openURL(menu)
		if err != nil {
			outputLabel.SetText(err.Error())
			return
		}

		// for a range ?
		if checkbox.Checked {
			// TODO: need to externalise cfg
			// Define the start date (08/01/2022)
			//startDate := time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC)
			startDate := date
			outputLabel.SetText("loop on week")
			// Get the current date
			currentDate := time.Now()

			// Iterate from the start date until the current date
			for date := startDate; date.Before(currentDate); date = date.AddDate(0, 0, 1) {
				// Check if the day is a Monday (Monday is the 1st day of the week in Go)
				if date.Weekday() == time.Monday {
					// Call the function to generate the menu for the week
					outputLabel.SetText(fmt.Sprintf("Starting process for date: %s", date.Format("02/01/2006")))
					generateWeekMenu(date, token)
				}
			}
			// end loop
		}
		myApp.Quit()
	})

	openButton := widget.NewButton("Open", func() {
		//TODO: get current week &/or date
		//url := "file:///Users/val/Documents/crockilo/crocmenu/menu_2023_18.html"
		path, err := getHTMLpath()

		pathrecipe := path + "/recipes"
		outputPath := pathrecipe + "/recettes.html"
		url := "file://" + outputPath

		outputLabel.SetText("opening " + url)
		err = openURL(url)
		if err != nil {
			outputLabel.SetText(err.Error())
			return
		}
	})

	indexButton := widget.NewButton("Index", func() {
		outputLabel.SetText("reindex")
		path, _ := getHTMLpath()

		pathrecipe := path + "/recipes"
		outputPath := pathrecipe + "/recettes.html"
		tmplPath := path + "/files.html"
		err := IndexMenu(pathrecipe, outputPath, tmplPath)
		if err != nil {
			outputLabel.SetText(err.Error())
		}
		outputLabel.SetText("reindex done")
		return
	})

	quitButton := widget.NewButton("Quit", func() {
		myApp.Quit()
	})

	myWin.SetContent(container.NewVBox(
		//container.NewHBox(dateEntry, checkbox),
		dateEntry,
		checkbox,
		//layout.NewSpacer(),
		container.NewHBox(
			startButton,
			openButton,
			indexButton,
			//layout.NewSpacer(),
			quitButton,
		),
		//layout.NewSpacer(),
		outputLabel,
	))

	//myWin.Resize(fyne.NewSize(400, 300))
	myWin.ShowAndRun()

}
