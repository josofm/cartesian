package mock

import (
	"github.com/josofm/cartesian/coordinate"
)

type CoordinateMock struct {
	points []coordinate.Point
}

func (c *CoordinateMock) CalculateRoute(vars map[string]string) []coordinate.Point {
	return c.points
}
