package main

import (
	"fmt"
	"os"
)

// certificate_monitor - Monitor SSL/TLS certs
func certificate_monitor(path string) {
	fmt.Println("========================================")
	fmt.Println("  Certificate-Monitor")
	fmt.Println("  Monitor SSL/TLS certs")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	certificate_monitor(path)
}
