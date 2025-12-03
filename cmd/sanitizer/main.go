package main

import (
	"fmt"
	"clipboard-sanitizer/internal/monitor"
)

func main() {
	fmt.Println("Clipboard sanitizer running...")
	monitor.Start()
}
