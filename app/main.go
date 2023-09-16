package main

import (
	"log"
	"net/http"
	"web-app-postgres/router"

	"github.com/rs/cors"
)

func main() {
	r := router.Router()

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	log.Fatal(http.ListenAndServe(":8080", corsWrapper.Handler(r)))
}
