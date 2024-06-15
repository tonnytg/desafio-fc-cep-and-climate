package domain

import (
	"fmt"
	"github.com/tonnytg/desafio-fc-cep-and-climate/internal/infra/cep"
	"github.com/tonnytg/desafio-fc-cep-and-climate/internal/infra/weather"
	"log"
)

type LocationService struct {
	repo LocationRepositoryInterface
}

type LocationServiceInterface interface{}

func NewLocationService(repo LocationRepositoryInterface) *LocationService {
	return &LocationService{
		repo: repo,
	}
}

func (s *LocationService) Execute(l *Location) error {

	//data := s.repo.Get(l.CEP)
	//log.Println("service received:", data)

	city, err := cep.GetCity(l.GetCEP())
	if err != nil {
		log.Println("error to get cep:", l.GetCEP())
		return fmt.Errorf("404")
	}

	if city == "" {
		log.Println("error to get cep:", l.GetCEP())
		return fmt.Errorf("404")
	}

	_ = l.SetCity(city)

	wc, err := weather.GetWeather(l.GetCity())
	if err != nil {
		log.Println("error to execute and get weather for city:", city)
		return fmt.Errorf("500")
	}
	_ = l.SetTemperatures(wc)

	err = s.repo.Save(l)
	if err != nil {
		log.Printf("error to save location: %v\n", l)
	}

	return nil
}
