package domain

import "errors"

var (
	ErrLocationNotFound = errors.New("Location not found")
)

type Location struct {
	CEP        string
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

type LocationInterface interface {
	GetCEP(cep string) (Location, error)
	SetCEP(cep string) error
	GetCelsius(cep string) float64
	GetFahrenheit(cep string) float64
	GetKelvin(cep string) float64
	SetCelsius(cep string, celsius float64) error
	SetFahrenheit(cep string, fahrenheit float64) error
	SetKelvin(cep string, kelvin float64) error
}

func NewLocation(cep string) (*Location, error) {

	if len(cep) != 8 {
		return nil, ErrLocationNotFound
	}

	return &Location{
		CEP: cep,
	}, nil
}

func (l *Location) GetCEP(cep string) (Location, error) {
	if l.CEP != cep {
		return Location{}, ErrLocationNotFound
	}

	return *l, nil
}

func (l *Location) SetCEP(cep string) error {
	l.CEP = cep
	return nil
}

func (l *Location) GetCelsius(cep string) float64 {
	return l.Celsius
}

func (l *Location) GetFahrenheit(cep string) float64 {
	return l.Fahrenheit
}

func (l *Location) GetKelvin(cep string) float64 {
	return l.Kelvin
}

func (l *Location) SetCelsius(cep string, celsius float64) error {
	l.Celsius = celsius
	return nil
}

func (l *Location) SetFahrenheit(cep string, fahrenheit float64) error {
	l.Fahrenheit = fahrenheit
	return nil
}

func (l *Location) SetKelvin(cep string, kelvin float64) error {
	l.Kelvin = kelvin
	return nil
}
