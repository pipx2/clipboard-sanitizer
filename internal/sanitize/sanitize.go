package sanitize

import "strings"

// Trims spaces
func Whitespace(text string) string {
	return strings.TrimSpace(text)
}


// Redact removes all letters and numbers as an example
func Redact(text string) string {
	return strings.Repeat(" ", len(text))  // simple redaction
}

// Mask replaces characters with asterisks except last 4
func Mask(text string) string {
	if len(text) <= 4 {
		return strings.Repeat("*", len(text))
	}
	mask := strings.Repeat("*", len(text)-4)
	return mask + text[len(text)-4:]
}
