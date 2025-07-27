// cmd/modernsynthetic/main.go
package main

import (
	"flag"
	"log"
	"os"

	"modernsynthetic/internal/modernsynthetic"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := modernsynthetic.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
