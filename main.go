package main

import (
	"github.com/tonnytg/desafio-fc-cep-and-climate/internal/domain"
	"log"
)

func main() {
	//webserver.Start()

	r := domain.NewLocationRepository()
	s := domain.NewLocationService(r)

	l, err := domain.NewLocation("12345678")
	if err != nil {
		log.Println("error to build location")
	}

	log.Println("TempC:", l.GetTempC())
	log.Println("TempF:", l.GetTempF())
	log.Println("TempK:", l.GetTempK())

	s.Execute(l)
}
