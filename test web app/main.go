package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(800, 500))
	myApp.Settings().SetTheme(theme.DarkTheme())

	entry := widget.NewMultiLineEntry()
	entry.Resize(fyne.NewSize(600, 300))
	entry.Move(fyne.NewPos(100, 135))

	button := widget.NewButton("Open file", func() {
		dialog.ShowFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				data, _ := io.ReadAll(reader)
				entry.SetText(string(data))
			},
			window,
		)
	})
	button.Resize(fyne.NewSize(150, 75))
	button.Move(fyne.NewPos(325, 30))

	cont := container.NewWithoutLayout(
		button,
		entry,
	)

	window.SetContent(cont)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun() // !!! change to show6 now ShowAndRun
}
