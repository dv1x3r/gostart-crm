package main

import (
	"log"

	"gostart-crm/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	a.MustRun()
}
