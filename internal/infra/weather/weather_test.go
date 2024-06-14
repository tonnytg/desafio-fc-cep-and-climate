package weather_test

import (
	"github.com/tonnytg/desafio-fc-cep-and-climate/internal/infra/weather"
	"os"
	"strings"
	"testing"
)

func TestWeatherGet(t *testing.T) {

	b, err := os.ReadFile("../../../.env")
	if err != nil {
		t.Errorf("file .env not found")
	}

	env := string(b)
	list := strings.Split(env, "=")
	if list[0] == "WEATHER_API_KEY" {
		_ = os.Setenv(list[0], list[1])
	}

	wc, err := weather.GetWeather("São Paulo")
	if err != nil {
		t.Error("error to get weather")
	}

	if wc < 1 {
		t.Error("sorry but something wrong with São Paulo")
	}

}
