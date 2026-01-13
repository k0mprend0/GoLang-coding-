package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func warpLabel(text string) *widget.Label {
	label := widget.NewLabel(text)
	label.Wrapping = fyne.TextWrapBreak
	return label
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(400, 400))
	myApp.Settings().SetTheme(theme.DarkTheme())

	img, _ := fyne.LoadResourceFromPath("IMG_9368.jpeg")

	cont1 := container.NewVBox(
		widget.NewLabel("Alex"),
		widget.NewEntry(),
		widget.NewButton("Click", func() {}),
	)

	cont2 := container.NewVBox(
		widget.NewLabel("Bro"),
		widget.NewEntry(),
		widget.NewButton("Click", func() {}),
	)

	cont3 := container.NewVBox(
		widget.NewLabel("Kate"),
		widget.NewEntry(),
		widget.NewButton("Click", func() {}),
	)

	cont4 := container.NewVBox(
		widget.NewLabel("John"),
		widget.NewEntry(),
		widget.NewButton("Click", func() {}),
	)

	// simple example
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Tab 1", img, cont1),
		container.NewTabItemWithIcon("Tab 2", img, cont2),
		container.NewTabItemWithIcon("Tab 3", img, cont3),
		container.NewTabItemWithIcon("Tab 4", img, cont4),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	window.SetContent(tabs)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun() // !!! change to show6 now ShowAndRun
	//window.SetMaster() // makes current window main
}
