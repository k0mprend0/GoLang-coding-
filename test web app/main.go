package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(900, 600))

	color_for := color.NRGBA{R: 136, G: 0, B: 147, A: 255}
	label := canvas.NewText("Color", color_for)

	window.SetContent(label)

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
