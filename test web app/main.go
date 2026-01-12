package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(400, 400))
	myApp.Settings().SetTheme(theme.DarkTheme())

	//автоматический показ/скрытие по нажатию кнопки
	//visibility := true

	label := widget.NewLabel("Some text here...")
	/*
			button := widget.NewButton("Change visibility", func() {
				if visibility {
					label.Hide()
				} else {
					label.Show()
				}

				visibility = !visibility
			})


		check := widget.NewCheck("Hide", func(visibility bool) {
			if visibility {
				label.Hide()
			} else {
				label.Show()
			}
		})
	*/

	window.SetContent(
		container.NewVBox(
			label,
			//button,
			//check,
		),
	)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
