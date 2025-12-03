package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/pipx2/clipboard-sanitizer/internal/sanitize"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
)

func main() {
	a := app.NewWithID("clipboard-sanitizer")
	w := a.NewWindow("Clipboard Sanitizer")

	w.Resize(fyne.NewSize(300, 150))

	infoLabel := widget.NewLabel("Press 'Sanitize Now' to clean clipboard.")
	sanitizeBtn := widget.NewButton("Sanitize Now", func() {
		text, err := clipboard.ReadAll()
		if err != nil {
			infoLabel.SetText("Error reading clipboard: " + err.Error())
			return
		}

		cleaned := sanitize.Run(text)
		if err := clipboard.WriteAll(cleaned); err != nil {
			infoLabel.SetText("Error writing clipboard: " + err.Error())
			return
		}

		infoLabel.SetText("Clipboard sanitized:\n" + cleaned)
		fmt.Println("Sanitized:", cleaned)
	})

	w.SetContent(container.NewVBox(
		infoLabel,
		sanitizeBtn,
	))

	// Show the window and start the app
	w.ShowAndRun()
}

