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

	/*
		window2 := myApp.NewWindow("window 2")
		window2.Resize(fyne.NewSize(400, 400))

		button := widget.NewButton("Open new window", func() {
			window2.Show()
			window2.SetContent(
				widget.NewLabel("Nostrud ullamco adipisicing in ea ad ut proident mollit."),
			)
		}
	*/

	window2 := myApp.NewWindow("Window 2")
	window2.Resize(fyne.NewSize(400, 400))

	entry := widget.NewMultiLineEntry()
	button := widget.NewButton("Show the text", func() {
		text := widget.NewLabel(entry.Text)
		text.Wrapping = fyne.TextWrapBreak
		window2.SetContent(
			container.NewVScroll(
				text,
			),
		)
		window2.Show()
	})

	close2 := widget.NewButton(
		"Close 2nd window",
		func() {
			window2.Close() // destroy window
			//window2.Hide() // hide window from user
		},
	)

	window.SetContent(
		container.NewVBox(
			entry,
			button,
			close2,
		),
	)
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.Show()      // !!! change to show6 now ShowAndRun
	window.SetMaster() // makes current window main
	myApp.Run()
}
