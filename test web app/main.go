package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(400, 400))
	myApp.Settings().SetTheme(theme.DarkTheme())

	names := []string{"Камень", "Ножницы", "Бумага"}

	list := widget.NewList(
		func() int { return len(names) },
		func() fyne.CanvasObject {
			return widget.NewButton("Press it", func() { fmt.Println("...") })
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Button).SetText(names[id])
		},
	)

	window.SetContent(list)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
