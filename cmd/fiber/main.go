package main

import (
	"log"

	"github.com/aadejanovs/catalog/internal/app"
)

func main() {
	server := app.Setup()
	log.Fatal(server.Listen("0.0.0.0:80"))
}
