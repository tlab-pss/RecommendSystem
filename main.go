package main

import (
	"github.com/joho/godotenv"
	"github.com/yuuis/RecommendSystem/api"
	"github.com/yuuis/RecommendSystem/infrastructures"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	s := infrastructures.NewServer()
	api.Router(s)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
