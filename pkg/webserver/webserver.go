package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/tonnytg/desafio-fc-cep-and-climate/internal/domain"
	"log"
	"net/http"
	"os"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func ReplyRequest(w http.ResponseWriter, statusCode int, msg string) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	replyMessage := ErrorMessage{Message: msg}

	err := json.NewEncoder(w).Encode(replyMessage)
	if err != nil {
		w.WriteHeader(statusCode)
		log.Println("")
		return fmt.Errorf("error to try reply request")
	}

	return nil
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {

	var data struct {
		CEP string `json:"cep"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		_ = ReplyRequest(w, http.StatusBadRequest, "no zipcode provided")
		return
	}

	log.Println("CEP API:", data.CEP)

	l, err := domain.NewLocation(data.CEP)
	if err != nil {
		log.Println(err)
		_ = ReplyRequest(w, http.StatusBadRequest, "no zipcode provided")
		return
	}

	repo := domain.NewLocationRepository()
	serv := domain.NewLocationService(repo)

	serv.Execute(l)

	responseData := struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
		TempK float64 `json:"temp_k"`
	}{
		l.GetTempC(),
		l.GetTempF(),
		l.GetTempK(),
	}

	byteResponseData, err := json.Marshal(responseData)
	if err != nil {
		_ = ReplyRequest(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteResponseData)

	return
}

func Start() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlerIndex)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Start webserver listen in port:", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Panicf("error to start http server")
	}
}
