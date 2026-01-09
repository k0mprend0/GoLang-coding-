package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(400, 400))
	myApp.Settings().SetTheme(theme.DarkTheme())

	file_item1 := fyne.NewMenuItem("New file", func() {
		os.Create("created.txt")
	})

	file_item2 := fyne.NewMenuItem("Save", func() {
		fmt.Println("File saved")
	})

	menu1 := fyne.NewMenu("File", file_item1, file_item2)

	actions_item1 := fyne.NewMenuItem("Hello", func() { fmt.Println("Hello!") })
	actions_item2 := fyne.NewMenuItem("Bye!", func() { fmt.Println("Goodbye!") })
	actions_item3 := fyne.NewMenuItem("Button", func() { fmt.Println("Clicked!") })

	menu2 := fyne.NewMenu("Actions", actions_item1, actions_item2, actions_item3)

	info_item1 := fyne.NewMenuItem("Info", func() { fmt.Println("Info!") })
	info_item2 := fyne.NewMenuItem("Clock me!", func() { fmt.Println("Info 'bout ur amazing app!") })
	info_item3 := fyne.NewMenuItem("Print", func() { fmt.Println("Printed!") })
	info_item4 := fyne.NewMenuItem("LOL", func() { fmt.Println("LOL!") })
	info_item5 := fyne.NewMenuItem("Menu item", func() { fmt.Println("Menu item") })
	info_item6 := fyne.NewMenuItem("xbox sucks", func() { fmt.Println("Thats true!") })

	menu3 := fyne.NewMenu(
		"Usless",
		info_item1,
		info_item2,
		info_item3,
		info_item4,
		info_item5,
		info_item6,
	)

	//всплывающие при наведении окна меню
	item1 := fyne.NewMenuItem("Options", nil)
	item2 := fyne.NewMenuItem("Say", nil)

	item1.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Print", func() {
			fmt.Println("Printed")
		}),
		fyne.NewMenuItem("Save", func() {
			fmt.Println("Saved")
		}),
		fyne.NewMenuItem("Cut", func() {
			fmt.Println("Cutted")
		}),
		fyne.NewMenuItem("Copy", func() {
			fmt.Println("Copied")
		}),
		fyne.NewMenuItem("Delete", func() {
			fmt.Println("Deleted")
		}),
	)

	item2.ChildMenu = fyne.NewMenu(
		"",
		fyne.NewMenuItem("Hi", func() {
			fmt.Println("Hello")
		}),
		fyne.NewMenuItem("Bye", func() {
			fmt.Println("Goodbye")
		}),
		fyne.NewMenuItem("LOL", func() {
			fmt.Println("Trolled")
		}),
	)

	menu4 := fyne.NewMenu("Buttons", item1, item2)

	main_menu := fyne.NewMainMenu(menu1, menu2, menu3, menu4)

	window.SetMainMenu(main_menu)

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}

