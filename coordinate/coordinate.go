package coordinate

import (
	"embed"
	"encoding/json"
	"log"
)

//go:embed data/points.json
var content embed.FS

type Point struct {
	X float64
	Y float64
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

func (c *Coordinate) CalculateRoute(vars map[string]string) []Point {
	return []Point{}
}
