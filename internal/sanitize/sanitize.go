package sanitize

import "strings"

func Run(text string) string {
	return strings.TrimSpace(text)
}
