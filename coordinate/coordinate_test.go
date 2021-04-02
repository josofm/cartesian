package coordinate_test

import (
	"testing"

	"github.com/josofm/cartesian/coordinate"
	"github.com/stretchr/testify/assert"
)

var params = []string{"x", "y", "distance"}

func TestShouldGetErrorConvertingValues(t *testing.T) {
	c := coordinate.NewCoordinate()

	vars := map[string]string{
		"key": "generic",
	}

	_, err := c.CalculateRoute(vars, params)
	assert.NotNil(t, err, "Should be null!")

}

func TestShoulGetPointsCorrectly(t *testing.T) {
	c := coordinate.NewCoordinate()
	lenExpected := 47

	vars := map[string]string{
		"x":        "-2",
		"y":        "-4",
		"distance": "100",
	}
	points, err := c.CalculateRoute(vars, params)

	assert.Equal(t, lenExpected, len(points))
	assert.Nil(t, err, "Should be null!")

}

func TestShoulGetErrorWhenDistanceIsNegative(t *testing.T) {
	c := coordinate.NewCoordinate()
	lenExpected := 0

	vars := map[string]string{
		"x":        "-2",
		"y":        "-4",
		"distance": "-1",
	}
	points, err := c.CalculateRoute(vars, params)

	assert.Equal(t, lenExpected, len(points))
	assert.NotNil(t, err, "Should be not null!")

}

func TestShoulGetErrorWhenXIsInvalid(t *testing.T) {
	c := coordinate.NewCoordinate()
	lenExpected := 0

	vars := map[string]string{
		"x":        "xis",
		"y":        "-4",
		"distance": "85",
	}
	points, err := c.CalculateRoute(vars, params)

	assert.Equal(t, lenExpected, len(points))
	assert.NotNil(t, err, "Should be not null!")
}

func TestShoulGetErrorWhenYIsInvalid(t *testing.T) {
	c := coordinate.NewCoordinate()
	lenExpected := 0

	vars := map[string]string{
		"x":        "-39",
		"y":        "not number bro",
		"distance": "85",
	}
	points, err := c.CalculateRoute(vars, params)

	assert.Equal(t, lenExpected, len(points))
	assert.NotNil(t, err, "Should be not null!")

}
