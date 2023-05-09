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
		menu := generateWeekMenu()
		err = openURL(menu)
		if err != nil {
			outputLabel.SetText(err.Error())
			return
		}
		myApp.Quit()
	})

	/*
		openButton := widget.NewButton("Open", func() {
			//TODO: get current week &/or date
			url := "file:///Users/val/Documents/crockilo/crocmenu/menu_2023_18.html"
			openURL(url)
		})
	*/
	quitButton := widget.NewButton("Quit", func() {
		myApp.Quit()
	})

	myWin.SetContent(container.NewVBox(
		dateEntry,
		//layout.NewSpacer(),
		container.NewHBox(
			startButton,
			//openButton,
			//layout.NewSpacer(),
			quitButton,
		),
		//layout.NewSpacer(),
		outputLabel,
	))

	//myWin.Resize(fyne.NewSize(400, 300))
	myWin.ShowAndRun()

}
