package router

import (
	"database/sql"
	"web-app-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", middleware.ProductsHandler).Methods("GET")
	router.HandleFunc("/info", middleware.NotImplementedYetHandler).Methods("GET")

	return router
}

func InitDB() *sql.DB {
	// create db connection
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost/productdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
