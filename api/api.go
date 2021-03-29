package api

import (
	"log"
	"net/http"
	"time"
)

type Api struct {
	server *http.Server
}

func NewApi() *Api {
	api := Api{}
	return &api
}

func (api *Api) StartServer() error {
	router := api.routes()
	muxWithMiddlewares := http.TimeoutHandler(router, time.Second*30, "Timeout!")
	api.server = &http.Server{
		Addr:    ":80",
		Handler: muxWithMiddlewares,
	}
	log.Print("Server is running at port 80")
	err := api.server.ListenAndServe()
	return err
}

func (api *Api) Up(w http.ResponseWriter, r *http.Request) {
	log.Print("[UP] Server is Up")
	w.WriteHeader(http.StatusOK)
}

func (api *Api) calculate(w http.ResponseWriter, r *http.Request) {
	log.Print("[calculate] method called")
	w.WriteHeader(http.StatusOK)
}
