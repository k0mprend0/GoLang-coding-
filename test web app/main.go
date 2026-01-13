package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(400, 400))
	//myApp.Settings().SetTheme(theme.DarkTheme())

	/*
		1 GridLayout
		2 GridWarpLayout
		3 CenterLayout
	*/

	/*
		// 1 GridLayout
		l1 := widget.NewLabel("Label 1")
		l2 := widget.NewLabel("Label 2")
		l3 := widget.NewLabel("Label 3")
		l4 := widget.NewLabel("Label 4")
		l5 := widget.NewLabel("Label 5")
		l6 := widget.NewLabel("Label 6")

		cont := container.New(
			layout.NewGridLayout(2),
			l1,
			l2,
			l3,
			l4,
			l5,
			l6,
		)
	*/

	/*
		// 2 GridWarpLayout
		l1 := widget.NewLabel("Label 1")
		l2 := widget.NewLabel("Label 2")
		l3 := widget.NewLabel("Label 3")
		l4 := widget.NewLabel("Label 4")
		l5 := widget.NewLabel("Label 5")
		l6 := widget.NewLabel("Label 6")

		cont := container.New(
			layout.NewGridWrapLayout(fyne.NewSize(100, 100)),
			l1,
			l2,
			l3,
			l4,
			l5,
			l6,
		)
	*/

	// 3 CenterLayout
	//img := canvas.NewImageFromFile("IMG_9368.jpeg")
	//img.FillMode = canvas.ImageFillOriginal
	//text := canvas.NewText("TEXT", color.RGBA{39, 255, 0, 255})

	label := widget.NewLabel("Label")
	entry := widget.NewEntry()
	button := widget.NewButton("Button", func() { fmt.Println(entry.Text) })

	cont_l := container.New(
		layout.NewCenterLayout(),
		label,
	)

	cont_e := container.New(
		layout.NewCenterLayout(),
		entry,
	)

	cont_b := container.New(
		layout.NewCenterLayout(),
		button,
	)

	window.SetContent(container.NewVBox(
		cont_l,
		cont_e,
		cont_b,
	))
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun() // !!! change to show6 now ShowAndRun
	//window.SetMaster() // makes current window main
}
