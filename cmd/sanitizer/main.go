package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/pipx2/clipboard-sanitizer/internal/monitor"
)
func main() {
	fmt.Println("Clipboard Sanitizer running... (Ctrl+C to stop)")

	// Stop gracefully on Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	done := make(chan bool)

	go func() {
		<-c
		fmt.Println("\nStopping sanitizer...")
		done <- true
	}()

	go monitor.Start() // start clipboard monitor

	<-done // wait for interrupt
	fmt.Println("Exited.")
}
