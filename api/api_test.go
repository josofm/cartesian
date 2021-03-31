package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/josofm/cartesian/api"
	"github.com/josofm/cartesian/mock"

	"github.com/stretchr/testify/assert"
)

type fixture struct {
	api *api.Api
	r   *mux.Router
}

func setup() fixture {
	c := &mock.CoordinateMock{}
	api := api.NewApi(c)

	router := mux.NewRouter()
	router.HandleFunc("/up", api.Up).Methods("GET")
	router.HandleFunc("/api/points/{x}/{y}/{distance}", api.Calculate).Methods("POST")

	return fixture{
		api: api,
		r:   router,
	}

}

func TestUpAPI(t *testing.T) {
	f := setup()

	r, err := http.NewRequest("GET", "/up", nil)

	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")

}

func TestShouldCalculateRouteCorreclty(t *testing.T) {
	f := setup()
	r, err := http.NewRequest("POST", "/api/points/8/2/2", nil)

	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusOK, rr.Code, "Status code Should be equal!")
}

func TestShouldGetBadRequestWhenParametersArentNumbers(t *testing.T) {
	f := setup()
	r, err := http.NewRequest("POST", "/api/points/this/not/numbers", nil)

	rr := httptest.NewRecorder()

	f.r.ServeHTTP(rr, r)

	assert.Nil(t, err, "Should be null!")
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Status code Should be equal!")
}
