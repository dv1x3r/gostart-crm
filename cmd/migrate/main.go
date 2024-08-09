package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/pressly/goose/v3"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/go-libsql"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]

	db, err := goose.OpenDBWithDriver(os.Getenv("DB_DRIVER"), os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.RunContext(context.Background(), command, db, "./migrations", arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
