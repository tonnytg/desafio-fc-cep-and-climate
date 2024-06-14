package domain

import (
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

func (s *LocationService) Execute(location *Location) {

	data := s.repo.Get(location.CEP)
	log.Println("service received:", data)

	err := s.repo.Save(location)
	if err != nil {
		log.Printf("error to save location: %v\n", location)
	}
}
