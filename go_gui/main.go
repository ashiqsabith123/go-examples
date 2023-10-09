package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Simple GUI App")

	// Create a button
	helloButton := widget.NewButton("Click Me!", func() {
		myWindow.SetContent(container.NewVBox(
			widget.NewLabel("Hello, World!"),
			widget.NewButton("Quit", func() {
				myApp.Quit()
			}),
		))
	})

	// Set the window's content to the button
	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Welcome to My Simple GUI App!"),
		helloButton,
	))

	// Show the window
	myWindow.ShowAndRun()
}
