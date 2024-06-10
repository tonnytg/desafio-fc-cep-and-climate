package main

import (
	"github.com/tonnytg/desafio-fc-cep-and-climate.git/internal/domain"
	"log"
)

func main() {

	repo := domain.NewLocationRepository()
	service := domain.NewLocationService(repo)

	useCase := domain.NewLocationUseCase(repo, service)
	location, err := useCase.Save("12345678")
	if err != nil {
		panic(err)
	}
	log.Printf("Location: %+v", location)

	location, err = useCase.Get("12345678")
	if err != nil {
		panic(err)
	}
	log.Printf("Location: %+v", location)
}
