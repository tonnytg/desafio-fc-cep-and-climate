package cep

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CEP struct {
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

func Get(cep string) (*CEP, error) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("error to get cep: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	var c CEP

	err = json.NewDecoder(response.Body).Decode(&c)
	if err != nil {
		log.Printf("error to decode cep: %s", err)
		return nil, err
	}

	if c.Cep == "" {
		return nil, fmt.Errorf("CEP not found")
	}

	return &c, nil
}
