package webserver

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	msg := ErrorMessage{Message: "success"}
	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("")
	}
	return
}

type ErrorMessage struct {
	Message string `json:"message"`
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
