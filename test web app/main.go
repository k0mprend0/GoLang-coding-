package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(400, 400))
	myApp.Settings().SetTheme(theme.DarkTheme())

	text := canvas.NewText("Text to color...", color.White)

	rec := canvas.NewRectangle(color.White)
	rec.SetMinSize(fyne.NewSize(300, 300))

	// dialog window ColorPicker
	cpt := dialog.NewColorPicker(
		"Color picker",
		"Choose your own color",
		func(c color.Color) {
			text.Color = c
			text.Refresh()

			rec.FillColor = c
			rec.Refresh()
		},
		window,
	)

	button := widget.NewButton("Color picker", func() {
		cpt.Show()
	})

	window.SetContent(
		container.NewVBox(
			button,
			text,
			rec,
		),
	)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun() // !!! change to show6 now ShowAndRun
}
