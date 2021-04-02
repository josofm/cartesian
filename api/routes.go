package api

import (
	"github.com/gorilla/mux"
)

func (api *Api) routes() *mux.Router {
	router := mux.NewRouter()
	router.Use()

	router.HandleFunc("/up", api.Up).Methods("GET")
	router.HandleFunc("/api/points/{x}/{y}/{distance}", api.calculate).Methods("GET")

	return router
}
