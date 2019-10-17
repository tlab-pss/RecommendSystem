package main

import (
	"log"
	"main/api"
	"main/infrastructures"
)

func main() {
	s := infrastructures.NewServer()
	api.Router(s)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
