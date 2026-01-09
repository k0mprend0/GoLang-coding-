package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Mini calculator")
	window.Resize(fyne.NewSize(300, 500))

	title := widget.NewLabel("Оформление заказа")

	name_label := widget.NewLabel("Ваше имя: ")
	name := widget.NewEntry()

	food_label := widget.NewLabel("Выберите еду для заказа: ")

	food := widget.NewCheckGroup([]string{"Pizza", "Cake", "Nuggets", "Burgers", "Soda"}, func(s []string) {})

	male_label := widget.NewLabel("Ваш пол: ")
	male := widget.NewRadioGroup([]string{"Муж", "Жен", "Оно"}, func(s string) {})

	result := widget.NewLabel("")

	button := widget.NewButton("Order it!", func() {
		username := name.Text
		food_arr := food.Selected
		male_val := male.Selected

		result.SetText(username + "\n" + male_val + "\n")

		for _, i := range food_arr {
			result.SetText(result.Text + i + "\n")
		}
	})

	window.SetContent(container.NewVBox(
		title,
		name_label,
		name,
		food_label,
		food,
		male_label,
		male,
		button,
		result,
	))

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
