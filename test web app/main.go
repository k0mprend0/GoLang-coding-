package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

	//progress := widget.NewProgressBar() // 0.0 - 1.0 by default
	//progress.Max = 100.0

	/*
		entry := widget.NewEntry()

		button := widget.NewButton("Change value", func() {
			value, _ := strconv.ParseFloat(entry.Text, 64)
			progress.SetValue(value)
			fmt.Println(progress.Value)
		})
	*/

	/*
		// example
		title := widget.NewLabel("Almost done!")

		name := widget.NewEntry()
		name.SetPlaceHolder("Enter your name")

		surname := widget.NewEntry()
		surname.SetPlaceHolder("Enter your surname")

		phone_number := widget.NewEntry()
		phone_number.SetPlaceHolder("Enter your phone number")

		nick := widget.NewEntry()
		nick.SetPlaceHolder("Enter your nick")

		password := widget.NewPasswordEntry()
		password.SetPlaceHolder("Your password")

		email := widget.NewEntry()
		email.SetPlaceHolder("Enter your email")

		bio := widget.NewMultiLineEntry()
		bio.SetPlaceHolder("Enter additional bio")

		tabs := container.NewAppTabs(
			container.NewTabItem(
				"Personal Info",
				container.NewVBox(
					name,
					surname,
					phone_number,
				),
			),
			container.NewTabItem(
				"Data for sign up",
				container.NewVBox(
					nick,
					password,
				),
			),
			container.NewTabItem(
				"Additional Info",
				container.NewVBox(
					email,
					bio,
				),
			),
		)

		next := widget.NewButton("Next step", func() {
			if tabs.SelectedIndex() < 2 {
				tabs.SelectIndex(tabs.SelectedIndex() + 1)
				progress.SetValue(progress.Value + 33)
			} else {
				progress.SetValue(progress.Value + 34)

				dialog.ShowInformation(
					"Sign up completed",
					"You have been signed up!",
					window,
				)
			}
		})

		window.SetContent(container.NewVBox(
			title,
			progress,
			tabs,
			next,
		))
	*/

	// progressbar infinite
	pbinfinite := widget.NewProgressBarInfinite()
	pbinfinite.Start()

	/*
		button := widget.NewButton("Click me", func() {
			if pbinfinite.Running() {
				pbinfinite.Stop()
			} else {
				pbinfinite.Start()
			}
		})
	*/

	// example
	pbinfinite.Hide()

	title := widget.NewLabel("Create your post")
	post_title := widget.NewEntry()
	post_title.SetPlaceHolder("your post title")

	post_text := widget.NewMultiLineEntry()
	post_text.SetPlaceHolder("Your text")

	submit := widget.NewButton("Submit", func() {
		pbinfinite.Show()
		time.Sleep(time.Second * 3)
		pbinfinite.Hide()

		dialog.ShowInformation(
			"Post creation",
			"Post have been created",
			window,
		)
	})

	window.SetContent(container.NewVBox(
		pbinfinite,
		title,
		post_title,
		post_text,
		submit,
	))
	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun() // !!! change to show6 now ShowAndRun
}
