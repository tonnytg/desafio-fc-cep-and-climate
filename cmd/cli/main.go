package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tonnytg/desafio-fc-cep-and-climate.git/internal/domain"
	"github.com/tonnytg/desafio-fc-cep-and-climate.git/internal/infra/cep"
	"github.com/tonnytg/desafio-fc-cep-and-climate.git/internal/infra/weather"
	"log"
)

type DTOTemperature struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func main() {

	location := domain.NewLocation("05541000")

	c, err := cep.Get(location.CEP)
	if err != nil {
		log.Println("error to get cep:", err)
		return
	}

	if c.Localidade == "" {
		log.Println("error localidade is empty")
		return
	}
	t, err := weather.Get(c.Localidade)
	if err != nil {
		log.Println("error to get weather:", err)
		return
	}

	tempC := t.Current.TempC
	tempF := float64(int((tempC*1.8+32)*10)) / 10
	tempK := float64(int((tempC+273.15)*10)) / 10

	var temperatures = DTOTemperature{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	byteTemp, err := json.Marshal(temperatures)
	if err != nil {
		log.Println("error to marshal json:", err)
		return
	}

	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, byteTemp, "", "\t")
	if err != nil {
		log.Println("error to indent json:", err)
		return
	}

	fmt.Println(string(prettyJSON.Bytes()))
}
