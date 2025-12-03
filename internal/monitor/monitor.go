package monitor

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/pipx2/clipboard-sanitizer/internal/sanitize"
)

func Start() {
	var last string
	pollInterval := 200 * time.Millisecond // 5 times per second

	for {
		current, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("Error reading clipboard:", err)
			time.Sleep(pollInterval)
			continue
		}

		if current != last {
			cleaned := sanitize.Run(current)
			if cleaned != current {
				clipboard.WriteAll(cleaned)
				fmt.Println("Sanitized clipboard:", cleaned)
			}
			last = cleaned
		}

		time.Sleep(pollInterval)
	}
}

