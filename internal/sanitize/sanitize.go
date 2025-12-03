package sanitize

import "strings"

// Run applies minimal sanitization (trim whitespace)
// Expand this function to add regex rules, redaction, masking, etc.
func Run(text string) string {
	return strings.TrimSpace(text)
}

