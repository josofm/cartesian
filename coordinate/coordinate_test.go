package coordinate_test

import (
	"sort"
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

	assert.Equal(t, lenExpected, len(points), "Should be equal!")
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

	assert.Equal(t, lenExpected, len(points), "Should be equal!")
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

	assert.Equal(t, lenExpected, len(points), "Should be equal!")
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

	assert.Equal(t, lenExpected, len(points), "Should be equal!")
	assert.NotNil(t, err, "Should be not null!")

}

func TestShouldValidateCorrectOrder(t *testing.T) {
	c := coordinate.NewCoordinate()
	lenExpected := 28

	vars := map[string]string{
		"x":        "-12",
		"y":        "-40",
		"distance": "75",
	}
	points, err := c.CalculateRoute(vars, params)

	assert.Equal(t, lenExpected, len(points), "Should be equal!")
	assert.True(
		t,
		sort.SliceIsSorted(points, func(i, j int) bool {
			return points[i].Distance < points[j].Distance
		}),
		"Should be true!")
	assert.Nil(t, err, "Should be not null!")
}
