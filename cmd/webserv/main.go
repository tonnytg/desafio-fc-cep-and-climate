package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/tonnytg/desafio-fc-cep-and-climate.git/internal/domain"
	"github.com/tonnytg/desafio-fc-cep-and-climate.git/internal/infra/cep"
	"github.com/tonnytg/desafio-fc-cep-and-climate.git/internal/infra/weather"
)

type ErroMessage struct {
	Message string `json:"message"`
}

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var data struct {
			CEP string `json:"cep"`
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			msg := ErroMessage{Message: "no zipcode provided"}
			json.NewEncoder(w).Encode(msg)
			return
		}

		location := domain.NewLocation(data.CEP)
		if location == nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			msg := ErroMessage{Message: "invalid zipcode"}
			json.NewEncoder(w).Encode(msg)
			return
		}

		log.Println("requesting weather for zipcode:", location.CEP)
		fullLocation, err := cep.Get(location.CEP)
		if err != nil {
			if err.Error() == "422" {
				w.WriteHeader(http.StatusUnprocessableEntity)
				msg := ErroMessage{Message: "invalid zipcode"}
				json.NewEncoder(w).Encode(msg)
				return
			}
			log.Println("3 Can not find zipcode")
			w.WriteHeader(http.StatusNotFound)
			msg := ErroMessage{Message: "can not find zipcode"}
			json.NewEncoder(w).Encode(msg)
			return
		}
		city := fullLocation.Localidade
		log.Println("requesting weather for city:", city)

		weatherLocation, err := weather.Get(city)
		if err != nil {
			log.Printf("error weather api can not find weather for city %s - %s\n", city, err)
			w.WriteHeader(http.StatusInternalServerError)
			msg := ErroMessage{Message: "can not find weather for city"}
			json.NewEncoder(w).Encode(msg)
			return
		}

		// format for 2 decimal places
		temperature := struct {
			TempC float64 `json:"temp_C"`
			TempF float64 `json:"temp_F"`
			TempK float64 `json:"temp_K"`
		}{
			TempC: weatherLocation.Current.TempC,
			TempF: float64((((weatherLocation.Current.TempC * 1.8) + 32) * 10) / 10),
			TempK: float64(((weatherLocation.Current.TempC + 273.15) * 10) / 10),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(temperature)
		return
	})

	log.Println("listening on port", PORT)
	if err := http.ListenAndServe(":"+PORT, mux); err != nil {
		panic(err)
	}
}
