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
	window.Resize(fyne.NewSize(250, 250))
	myApp.Settings().SetTheme(theme.DarkTheme())

	label := widget.NewLabel("Mini notepad")

	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("Type sth here...")
	//entry.Wrapping = fyne.TextWrapBreak //адаптивно к размерам окна
	//entry.Wrapping = fyne.TextWrapOff //размеры приложения адаптируются к длине строки

	button := widget.NewButton("Save", func() {
		file, err := os.Create("Info.txt")
		if err != nil {
			fmt.Println("Error! - ", err)
			os.Exit(1)
		}
		defer file.Close()

		file.WriteString(entry.Text)
	})

	content := container.NewVBox(label, entry, button)

	window.SetContent(content)

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
