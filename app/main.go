package main

import (
	"log"
	"net/http"
	"web-app-postgres/router"

	"github.com/rs/cors"
)

func main() {
	// add postgres db connection
	db := router.InitDB()
	defer db.Close()

	// add router
	r := router.Router()

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	log.Fatal(http.ListenAndServe(":8080", corsWrapper.Handler(r)))
}
