package main

import (
	"log"

	"github.com/josofm/cartesian/api"
	"github.com/josofm/cartesian/coordinate"
)

func main() {
	log.Print("[Main] Starting application!")

	c := coordinate.NewCoordinate()

	if err := api.NewApi(c).StartServer(); err != nil {
		panic("Panic starging server!")
	}

}
