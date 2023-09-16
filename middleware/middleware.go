package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"web-app-postgres/domain"
)

func NotImplementedYetHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal(w.Write([]byte("Not Implemented")))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(domain.GetSampleProducts())

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
