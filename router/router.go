package router

import (
	"web-app-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", middleware.ProductsHandler).Methods("GET")
	router.HandleFunc("/info", middleware.NotImplementedYetHandler).Methods("GET")

	return router
}
