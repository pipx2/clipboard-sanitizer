package monitor

import {
	"clipboard-sanitizer/internal/sanitize"
	"github.com/atotto/clipboard"
	"time"
}

func Start() {
	var last string

	for {
		current, _ := clipboard.ReadAll()

		if current != last {
			cleaned := sanitize.Run(current)
			
			if cleaned != current {
				clipboard.WriteAll(cleaned)
			}
			last = cleaned
		}
		time.Sleep(200 * time.Millisecond)
	}
}
