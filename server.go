package main

import (
	"gqlgen-lambda/routers"
	"log"
	"net/http"
)

const defaultPort = ":5000"

func main() {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:    defaultPort,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
