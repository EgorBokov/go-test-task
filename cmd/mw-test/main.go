package main

import (
	"log"
	"middleware/internal/pkg/app"
)

func main() {
	server, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
