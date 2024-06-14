package cep

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FullLocation struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func Get(cep string) (*FullLocation, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	// Create a custom HTTP client with a Transport that ignores SSL verification
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	response, err := httpClient.Get(url)
	if err != nil {
		log.Printf("error to get cep: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// can not find zipcode
		return nil, fmt.Errorf("404")
	}

	var f FullLocation

	err = json.NewDecoder(response.Body).Decode(&f)
	if err != nil || f.Cep == "" {
		// invalid zipcode
		return nil, fmt.Errorf("422")
	}

	return &f, nil
}
