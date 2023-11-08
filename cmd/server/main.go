package main

import (
	"log"
	"w2go/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	a.MustRun()
}
