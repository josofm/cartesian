package mock

import (
	"github.com/josofm/cartesian/coordinate"
)

type CoordinateMock struct {
	points []coordinate.Point
	Err    error
}

func (c *CoordinateMock) CalculateRoute(vars map[string]string, params []string) ([]coordinate.Point, error) {
	return c.points, c.Err
}
