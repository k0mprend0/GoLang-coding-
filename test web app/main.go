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
	window.Resize(fyne.NewSize(800, 500))
	myApp.Settings().SetTheme(theme.DarkTheme())

	names := []string{
		"Kate",
		"John",
		"Anrew",
	}

	marks := [][]int{
		{5, 3, 5},
		{4, 5, 2},
		{2, 4, 4},
	}

	table := widget.NewTable(
		func() (rows int, cols int) {
			return len(names), len(marks[0]) + 1
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Default text")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			if tci.Col == 0 {
				co.(*widget.Label).SetText(names[tci.Row])
			} else {
				co.(*widget.Label).SetText(fmt.Sprint(marks[tci.Row][tci.Col-1]))
			}
		},
	)

	window.SetContent(table)
	window.ShowAndRun() // !!! change to show6 now ShowAndRun
}
