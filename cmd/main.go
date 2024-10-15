package main

import (
	"log"

	"github.com/realiza-org/go-service/internal/server"
)

func main() {
	s := server.NewServer()
	log.Println("Starting server on :8080...")
	log.Fatal(s.Start(":8080"))
}
