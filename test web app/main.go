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

	save_file := widget.NewButton("Save file", func() {
		dialog.ShowFileSave(
			func(writer fyne.URIWriteCloser, err error) {
				io.WriteString(writer, entry.Text)
			},
			window,
		)
	})

	window.SetContent(
		container.NewVBox(
			entry,
			save_file,
		),
	)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun() // !!! change to show6 now ShowAndRun
}
