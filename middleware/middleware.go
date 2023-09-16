package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"web-app-postgres/domain/product"
)

func NotImplementedYetHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal(w.Write([]byte("Not Implemented")))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(product.GetSampleProducts())

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
