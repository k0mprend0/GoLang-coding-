package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("My app")
	window.Resize(fyne.NewSize(400, 400))
	myApp.Settings().SetTheme(theme.DarkTheme())

	card := widget.NewCard(
		"VIMRT",
		"Some info about esport team",
		widget.NewAccordion(
			widget.NewAccordionItem(
				"Teams roster",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Alexey k0mprend0 Popov",
						widget.NewLabel("Info about player"),
					),
					widget.NewAccordionItem(
						"Kirill Schaidt Podrez",
						widget.NewLabel("Info about player"),
					),
					widget.NewAccordionItem(
						"Artem artemixx Veselov",
						widget.NewLabel("Info about player"),
					),
				),
			),
			widget.NewAccordionItem(
				"Teams achivements",
				widget.NewAccordion(
					widget.NewAccordionItem(
						"Faceit legue season 5",
						widget.NewLabel("Info about tournament"),
					),
					widget.NewAccordionItem(
						"Stealseries i-legue",
						widget.NewLabel("Info about tournament"),
					),
					widget.NewAccordionItem(
						"Blast Moscow",
						widget.NewLabel("Info about tournament"),
					),
				),
			),

			widget.NewAccordionItem(
				"More info",
				widget.NewLabel("No info yet"),
			),
		),
	)

	window.SetContent(card)

	/*
		text := "puhasdfihlgadfhasoufhsiuodhfgladhnfojkladbnfilhadgivojhasgiuoadgfoyusdatfiouahnmfoetriunvdgoiygdioyfgvnoidyfgnvioydghnmf8yovtoyiufegnv8ywg 97we9f vw97 rt97wew90puqhg97pwegpiu g8owe w9peuy7"
		text_label := widget.NewLabel(text)
		text_label.Wrapping = fyne.TextWrapBreak
		item1 := widget.NewAccordionItem(
			"Read more info",
			text_label,
		)

		item2 := widget.NewAccordionItem(
			"Button",
			widget.NewButton(
				"Say Hello!",
				func() {
					fmt.Println("Hello!")
				},
			),
		)

		item3 := widget.NewAccordionItem(
			"Chapters",
			widget.NewAccordion(
				widget.NewAccordionItem(
					"Chapter 1",
					widget.NewLabel("Some text here..."),
				),
				widget.NewAccordionItem(
					"Chapter 2",
					widget.NewLabel("Some text from chapter 2 here..."),
				),
				widget.NewAccordionItem(
					"Chapter 3",
					widget.NewLabel("Very interesting text here..."),
				),
			),
		)

		accordion := widget.NewAccordion(item1, item2, item3)

		window.SetContent(accordion)
	*/

	//window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}
