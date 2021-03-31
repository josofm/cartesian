package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/josofm/cartesian/coordinate"
)

const (
	errorParams = "Params must be valid numbers!"
)

type Api struct {
	server *http.Server
	c      Coordinate
}

var params = []string{"x", "y", "distance"}

type Coordinate interface {
	CalculateRoute(vars map[string]string) []coordinate.Point
}

func NewApi(c Coordinate) *Api {
	api := Api{
		c: c,
	}
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
	vars := mux.Vars(r)
	if !validateParams(vars) {
		log.Print("[calculate] invalid params!")
		sendErrorMessage(w, http.StatusBadRequest, errorParams)
		return
	}

	distances := api.c.CalculateRoute(vars)
	log.Println("[calculate] calculate ok - ", distances)
	send(w, http.StatusOK, distances)
	return

}

func validateParams(vars map[string]string) bool {
	for _, p := range params {
		if _, ok := vars[p]; !ok {
			return false
		}
		if _, err := strconv.ParseFloat(vars[p], 64); err != nil {
			return false
		}
	}
	return true
}

func sendErrorMessage(w http.ResponseWriter, code int, msg string) {
	log.Printf("Error - %s", msg)
	send(w, code, msg)
}

func send(w http.ResponseWriter, code int, val interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if val != nil {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			log.Printf("error on json encoder err: %s", err.Error())
		}
	}
}
