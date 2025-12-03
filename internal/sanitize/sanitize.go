package sanitize

import "strings"

// Minimal sanitizer: trims spaces
func Run(text string) string {
	return strings.TrimSpace(text)
}

