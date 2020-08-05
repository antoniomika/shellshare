// Package main represents the main entrypoint of the shellshare application.
package main

import (
	"log"

	"github.com/antoniomika/shellshare/cmd"
)

// main will start the shellshare command lifecycle and spawn the shellshare services.
func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println("Unable to execute root command:", err)
	}
}
