package main

import (
	"github.com/tonnytg/desafio-fc-cep-and-climate/internal/domain"
	"github.com/tonnytg/desafio-fc-cep-and-climate/internal/infra/cep"
	"github.com/tonnytg/desafio-fc-cep-and-climate/internal/infra/weather"
	"log"
)

func main() {
	//webserver.Start()

	r := domain.NewLocationRepository()
	s := domain.NewLocationService(r)

	l, err := domain.NewLocation("05541000")
	if err != nil {
		log.Println("error to build location")
	}

	city, _ := cep.GetCity(l.GetCEP())
	l.SetCity(city)

	wc, _ := weather.GetWeather(l.GetCity())
	_ = l.SetTemperatures(wc)

	log.Println("TempC:", l.GetTempC())
	log.Println("TempF:", l.GetTempF())
	log.Println("TempK:", l.GetTempK())

	s.Execute(l)

	//city, err := cep.GetCity("05541000")
	//if err != nil {
	//	log.Println("error to get cep:", err)
	//}
	//
	//w, err := weather.GetWeather(city)
	//if err != nil {
	//	log.Println("error to get weather:", err)
	//}
	//log.Println("weather:", w)
}
