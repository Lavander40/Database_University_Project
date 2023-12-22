package main

import (
	"postgre_api/server"
	"log"
)

func main() {
	s := server.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}