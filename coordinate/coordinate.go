package coordinate

import (
	"embed"
	"encoding/json"
	"errors"
	"log"
	"strconv"
)

//go:embed data/points.json
var content embed.FS

type Point struct {
	X int
	Y int
}

type Coordinate struct {
	Points []Point
}

//Since this api needs a file with points to works fine
//if we can't read the file we panic
//in another context, we would treat error
func NewCoordinate() *Coordinate {
	log.Print("Create a new coordinate")
	raw, err := content.ReadFile("data/points.json")
	if err != nil {
		panic("Can't read points file!")
	}
	var p []Point

	if err := json.Unmarshal(raw, &p); err != nil {
		panic("Can't decode the file!")
	}

	return &Coordinate{
		Points: p,
	}

}

func (c *Coordinate) CalculateRoute(vars map[string]string, params []string) ([]Point, error) {
	x, y, distance, err := convertValues(vars, params)
	if err != nil {
		log.Print("[CalculateRoute] error convert values")
		return []Point{}, err
	}
	p := Point{
		X: x,
		Y: y,
	}
	orderedPoints := c.orderPoints(distance, p)
	log.Print("[CalculateRoute] Calculated route succefully")
	return orderedPoints, nil
}

func (c *Coordinate) orderPoints(distance int, base Point) []Point {
	var result []Point
	for _, p := range c.Points {
		if abs((base.X-p.X))+abs((base.Y-p.Y)) <= distance {
			result = append(result, p)
		}
	}

	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func convertValues(vars map[string]string, params []string) (int, int, int, error) {
	x, err := strconv.Atoi(vars[params[0]])
	if err != nil {
		return 0, 0, 0, err
	}
	y, err := strconv.Atoi(vars[params[1]])
	if err != nil {
		return 0, 0, 0, err
	}
	distance, err := strconv.Atoi(vars[params[2]])
	if err != nil {
		return 0, 0, 0, err
	}
	if distance < 0 {
		log.Print("[CalculateRoute] distance must be a positive number")
		return 0, 0, 0, errors.New("Distance must be a positive number")
	}
	return x, y, distance, nil
}
