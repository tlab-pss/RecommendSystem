package main

import (
	"log"

	"github.com/yuuis/RecommendSystem/api"
	"github.com/yuuis/RecommendSystem/infrastructures"
)

func main() {
	infrastructures.InitEnvironment()

	s := infrastructures.NewServer()
	api.Router(s)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
