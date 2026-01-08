package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Mini calculator")
	window.Resize(fyne.NewSize(400, 320))

	label1 := widget.NewLabel("Enter ur 1st number")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Enter ur 2nd number")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("")

	button := widget.NewButton("Start calculating", func() {
		n1, err1 := strconv.ParseFloat(entry1.Text, 64) // конвертация из str в float (64 - колво бит)
		n2, err2 := strconv.ParseFloat(entry2.Text, 64)

		if err1 != nil || err2 != nil {
			answer.SetText("Input error!")
		} else {
			sum := n1 + n2
			sub := n1 - n2
			mult := n1 * n2
			div := n1 / n2

			answer.SetText(fmt.Sprintf("Сумма: %f\nРазность: %f\nПроизведение: %f\nЧастное: %f", sum, sub, mult, div))
		}
	})

	window.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		button,
		answer,
	))

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
