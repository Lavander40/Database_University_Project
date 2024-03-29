package main

import (
	"elastic_api/server"
	"log"
)

func main() {
	s := server.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}