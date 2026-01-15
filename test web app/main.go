package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(800, 500))
	myApp.Settings().SetTheme(theme.DarkTheme())

	username := widget.NewEntry()
	email := widget.NewEntry()
	password := widget.NewPasswordEntry()

	form := widget.NewForm(
		widget.NewFormItem("Username", username),
		widget.NewFormItem("Email", email),
		widget.NewFormItem("Password", password),
	)

	form.OnSubmit = func() {
		log.Println("Form submited!")

		log.Printf(
			"Username: %s\nEmail: %s\nPassword: %s\n",
			username.Text,
			email.Text,
			password.Text,
		)
	}

	form.SubmitText = "Send"

	form.OnCancel = func() {
		myApp.Quit()
	}

	form.CancelText = "Quit"

	window.SetContent(
		container.NewVBox(
			form,
		),
	)
	window.ShowAndRun()
}
