package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Hello World")

	label := widget.NewLabel("Hello everyone!")

	//input field
	entry := widget.NewEntry()

	//creating buttons
	button := widget.NewButton("Press me!", func() {
		data := entry.Text // получение данных из input fields
		fmt.Println(data)
	})

	window.SetContent(container.NewVBox(
		label,
		entry,
		button,
	))

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
