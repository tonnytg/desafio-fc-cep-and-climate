package domain

import "fmt"

type Location struct {
	CEP   string
	City  string
	TempC float64
	TempF float64
	TempK float64
}

func NewLocation(cep string) (*Location, error) {

	var l Location

	err := l.SetCEP(cep)
	if err != nil {
		return nil, fmt.Errorf("error to create object Location")
	}

	err = l.SetTemperatures(0)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (l *Location) GetCEP() string {
	return l.CEP
}

func (l *Location) GetTempC() float64 {
	return l.TempC
}

func (l *Location) GetTempF() float64 {
	return l.TempF
}

func (l *Location) GetTempK() float64 {
	return l.TempK
}

func (l *Location) SetCEP(cep string) error {

	if len(cep) != 8 {
		return fmt.Errorf(" invalid zipcode")
	}

	l.CEP = cep
	return nil
}

func (l *Location) SetTempC(celcius float64) error {

	l.TempC = celcius

	return nil
}

func (l *Location) setTempF() error {
	l.TempF = (l.TempC * 1.8) + 32
	return nil
}

func (l *Location) setTempK() error {
	l.TempK = l.TempC + 273
	return nil
}

func (l *Location) SetTemperatures(celcius float64) error {

	err := l.SetTempC(celcius)
	if err != nil {
		return err
	}
	err = l.setTempF()
	if err != nil {
		return err
	}
	err = l.setTempK()
	if err != nil {
		return err
	}

	return nil
}
