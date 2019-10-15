package main

import (
	"main/api"
	"main/infrastructures"
	"log"
)

func main() {
	s := infrastructures.NewServer()
	api.Router(s)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
