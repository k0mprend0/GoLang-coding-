package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(400, 400))
	myApp.Settings().SetTheme(theme.DarkTheme())

	//types of dialog windows (not all, just some of 'em)
	//1 Information
	//2 Confirm
	//3 CustomConfirm
	//4 Error
	//5 Custom

	/*
		// 1 Information
		button := widget.NewButton("Click me", func() {
			dialog.ShowInformation(
				"Information about programm",
				"Quis anim nisi in velit veniam ad dolore elit reprehenderit aliqua.",
				window,
			)
		})
	*/

	/*
		// 2 Confirm
		button := widget.NewButton("Click me", func() {
			dialog.ShowConfirm(
				"Do you agree that Grisha dumbass?",
				"This is annonymus pool",
				func(b bool) {
					if b {
						fmt.Println("This is correct!")
					} else {
						fmt.Println("Incorrect")
					}
				},
				window,
			)
		})
	*/

	/*
		// 3 CustomConfirm
		button := widget.NewButton("Click me", func() {
			dialog.ShowCustomConfirm(
				"Grisha daun?",
				"YES!",
				"OF COURSE!",
				widget.NewLabel("Commodo veniam et Lorem voluptate."),
				func(b bool) {
					if b {
						fmt.Println("This is correct!")
					} else {
						fmt.Println("PERFECT")
					}
				},
				window,
			)
		})
	*/

	/*
		// 4 Error
		button := widget.NewButton("Click me", func() {
			dialog.ShowError(
				errors.New("Unknown issue"),
				window,
			)
		})
	*/

	// 5 Custom
	button := widget.NewButton("Click me", func() {
		dialog.ShowCustom(
			"Information for user",
			"I have read it",
			widget.NewLabel("Duis nulla eiusmod ullamco consequat ea."),
			window,
		)
	})

	window.SetContent(button)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
