package main

import (
	"fmt"

	"github.com/atotto/clipboard"  // Clipboard access
	"github.com/pipx2/clipboard-sanitizer/internal/sanitize"  // Sanitization functions
	"fyne.io/fyne/v2/app"  // Fyne app framework
	"fyne.io/fyne/v2/container"  // Containers for layout
	"fyne.io/fyne/v2/widget"  // Widgets (buttons, labels, radio buttons)
	"fyne.io/fyne/v2"  // Fyne core types
)

// main initializes the Clipboard Sanitizer GUI application.
// The app reads text from the system clipboard, applies a sanitization
// method (e.g. Redact, Mask and Whitespace).
// The GUI allows the user to select the mode and trigger sanitization
// via a button.
func main() {
	// Create a new Fyne application with a unique ID
	a := app.NewWithID("clipboard-sanitizer")

	// Create a new window for the application
	w := a.NewWindow("Clipboard Sanitizer")
	w.Resize(fyne.NewSize(300, 150))  // Set window size

	// Label to display instructions and status messages
	infoLabel := widget.NewLabel("Select mode and press 'Sanitize Now'.")

	// Radio buttons for selecting the sanitization mode
	modes := widget.NewRadioGroup([]string{"Redact", "Mask", "Whitespace"}, func(s string) {
		// Update label whenever the mode changes
		infoLabel.SetText(fmt.Sprintf("Mode selected: %s. Press 'Sanitize Now'.", s))
	})
	modes.Selected = "Redact"  // default selection

	// Button to trigger sanitization
	var cleaned string  // kept outside to share between callbacks
	sanitizeBtn := widget.NewButton("Sanitize Now", func() {
		// Read text from the system clipboard
		text, err := clipboard.ReadAll()
		if err != nil {
			infoLabel.SetText("Error reading clipboard: " + err.Error())
			return
		}

		// Apply the selected sanitization method
		switch modes.Selected {
		case "Redact":
			cleaned = sanitize.Redact(text)
		case "Mask":
			cleaned = sanitize.Mask(text)
		case "Whitespace":
			cleaned = sanitize.Whitespace(text)
		default:
			cleaned = text
		}


		// Update the label with confirmation and print to console
		infoLabel.SetText("Clipboard sanitized:\n" + cleaned)
		fmt.Println("Sanitized:", cleaned)
	})

	// Add button to move sanitized to clipboard
	updateClipboardBtn := widget.NewButton("Update Clipboard", func() {
		// Write the sanitized text back to the clipboard
		if err := clipboard.WriteAll(cleaned); err != nil {
			infoLabel.SetText("Error writing clipboard: " + err.Error())
			return
		}

		// If clipboard write was successful
		infoLabel.SetText("Clipboard updated successfully!")
	})

	// Arrange widgets vertically in the window
	w.SetContent(container.NewVBox(
		infoLabel,
		modes,
		sanitizeBtn,
		updateClipboardBtn,
	))

	// Show the window and start the application event loop
	w.ShowAndRun()
}

