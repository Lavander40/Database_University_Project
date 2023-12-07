package main

import (
	"db_project/internal/db_project/microservices/redis_api/server"
	"log"
)

func main() {
	s := server.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
