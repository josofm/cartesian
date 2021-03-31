package api

import "net/http"

func (api *Api) Calculate(w http.ResponseWriter, r *http.Request) {
	api.calculate(w, r)
}
