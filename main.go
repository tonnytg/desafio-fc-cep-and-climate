package main

import (
	"github.com/tonnytg/desafio-fc-cep-and-climate/pkg/webserver"
	"log"
)

func main() {

	log.Println("Start API Server")
	webserver.Start()
}
