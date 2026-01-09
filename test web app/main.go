package main

import (
	"fmt"
	"os"

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

	label := widget.NewLabel("Registration")
	label.TextStyle = fyne.TextStyle{Bold: true} //изменение стилей текста
	label.Move(fyne.NewPos(137, 3))

	username := widget.NewEntry()
	username.Resize(fyne.NewSize(300, 40))
	username.Move(fyne.NewPos(50, 50))
	username.SetPlaceHolder("Enter username")

	password := widget.NewPasswordEntry()
	password.Resize(fyne.NewSize(300, 40))
	password.Move(fyne.NewPos(50, 110))
	password.SetPlaceHolder("Enter password")

	email := widget.NewEntry()
	email.Resize(fyne.NewSize(300, 40))
	email.Move(fyne.NewPos(50, 170))
	email.SetPlaceHolder("Enter your email")

	info := widget.NewMultiLineEntry()
	info.Resize(fyne.NewSize(300, 80))
	info.Move(fyne.NewPos(50, 230))
	info.SetPlaceHolder("Enter more info")

	submit := widget.NewButton("Register", func() {
		file, err := os.Create("user.txt")

		if err != nil {
			fmt.Println("Error! - ", err)
			os.Exit(1)
		}
		defer file.Close()

		data := fmt.Sprintf(
			"Username: %s\nPassword: %s\nEmail: %s\nInfo: %s\n",
			username.Text,
			password.Text,
			email.Text,
			info.Text,
		)
		file.WriteString(data)
	})
	submit.Resize(fyne.NewSize(200, 40))
	submit.Move(fyne.NewPos(100, 340))

	cont := container.NewWithoutLayout(
		label,
		username,
		password,
		email,
		info,
		submit,
	)

	window.SetContent(cont)

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
