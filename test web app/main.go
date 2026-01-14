package main

import (
	"fmt"

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

	//label := widget.NewLabel("Slider value")

	slider := widget.NewSlider(0.0, 10.0)
	//slider.Value = 50.0 // set default value

	/*
		entry := widget.NewEntry()

		button := widget.NewButton("Set value", func() {
			value, _ := strconv.ParseFloat(entry.Text, 64)
			slider.Value = value
			slider.Refresh()
		})

		window.SetContent(container.NewVBox(
			slider,
			entry,
			button,
		))
	*/

	/*
		// OnChange method
		slider.OnChanged = func(f float64) {
			label.SetText(fmt.Sprintf("%f", f))
		}

		window.SetContent(container.NewVBox(
			label,
			slider,
		))
	*/

	title := widget.NewLabel("Evaluate the app's perfomance from 0 to 10")
	label := widget.NewLabel("Tell about your experience")

	feed := widget.NewLabel("Your score wiil be displayed here")

	entry := widget.NewEntry()
	entry.PlaceHolder = "Enter your feedback"

	button := widget.NewButton("Send feedback", func() { fmt.Println(entry.Text) })

	label.Hide()
	entry.Hide()
	button.Hide()

	slider.OnChanged = func(f float64) {
		feed.SetText("Your score: " + fmt.Sprintf("%f", f))

		if f < 5 {
			label.Show()
			label.SetText("Tell about your experience")
			entry.Show()
			button.Show()
		} else {
			label.SetText("Thanks for your feedback")
			entry.Hide()
			button.Hide()
		}
	}

	window.SetContent(container.NewVBox(
		title,
		slider,
		feed,
		label,
		entry,
		button,
	))

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun() // !!! change to show6 now ShowAndRun
	//window.SetMaster() // makes current window main
}
