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

	//выпадающий список с доп вводом от пользователя
	label := widget.NewLabel("")

	sel := widget.NewSelectEntry(
		[]string{
			"Option 1",
			"Option 2",
			"Option 3",
			"Option 4",
			"Option 5",
		},
	)

	sel.PlaceHolder = "Choose one or enter another option"

	button := widget.NewButton("Get data", func() {
		label.SetText("You peak " + sel.Text)
	})

	/*

		//обычный выпадаюзий список
		label := widget.NewLabel("")

		sel := widget.NewSelect(
			[]string{
				"Option 1",
				"Option 2",
				"Option 3",
				"Option 4",
				"Option 5",
			},
			nil,
		)

		sel.PlaceHolder = "Choose one option"

		button := widget.NewButton("Get data", func() {
			label.SetText("You peak " + sel.Selected)
		})
	*/
	window.SetContent(
		container.NewVBox(
			sel,
			button,
			label,
		),
	)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
