package domain

import (
	"fmt"
	"log"
)

type Location struct {
	CEP   string  `json:"cep"`
	City  string  `json:"city"`
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

func NewLocation(cep string) *Location {

	l := &Location{}

	err := l.SetCEP(cep)
	if err != nil {
		return nil
	}

	return l
}

func (l *Location) SetCEP(cep string) error {

	if len(cep) != 8 {
		log.Printf("CEP inválido")
		return fmt.Errorf("CEP inválido")
	}
	l.CEP = cep

	return nil
}
