// +build integration

package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/josofm/cartesian/api"
	"github.com/josofm/cartesian/coordinate"

	"github.com/stretchr/testify/assert"
)

const (
	baseUrl = "http://localhost:80"
)

func setupIntegration() {
	c := coordinate.NewCoordinate()

	a := api.NewApi(c)
	go a.StartServer()
	WaitServerUp()
}

func WaitServerUp() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for range ticker.C {
		res, err := http.Get(baseUrl + "/up")
		if err != nil {
			continue
		}
		if res.StatusCode == http.StatusOK {
			return
		}
	}
}

func TestShouldCalculatePointsCorrectly(t *testing.T) {
	setupIntegration()
	expectedSize := 83
	resp, err := http.Get(baseUrl + "/api/points/-80/9/200")

	resultActual := parseBody(resp.Body)
	defer resp.Body.Close()

	assert.Equal(t, expectedSize, len(resultActual), "Should be not nil!")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldCalculatePointsCorrectlyButNoPointAreReturned(t *testing.T) {
	setupIntegration()
	expectedSize := 0
	resp, err := http.Get(baseUrl + "/api/points/0/0/1")

	resultActual := parseBody(resp.Body)
	defer resp.Body.Close()

	assert.Equal(t, expectedSize, len(resultActual), "Should be not nil!")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldGetBadRequestNegativeDistance(t *testing.T) {
	setupIntegration()

	resp, err := http.Get(baseUrl + "/api/points/-8/-9/-2")

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

func TestShouldGetBadRequestInvalidFieldTypes(t *testing.T) {
	setupIntegration()

	resp, err := http.Get(baseUrl + "/api/points/this/is/word")

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Should be equal!")
	assert.Nil(t, err, "Should be nil!")
}

//skip errors because it isn't focus of the test
func parseBody(body io.Reader) []map[string]interface{} {
	decoder := json.NewDecoder(body)
	var bodyActual []map[string]interface{}
	_ = decoder.Decode(&bodyActual)
	return bodyActual
}
