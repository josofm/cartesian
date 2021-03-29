package main

import (
	"fmt"
	"log"

	"github.com/josofm/cartesian/api"
	"github.com/josofm/cartesian/coordinate"
)

func main() {
	log.Print("[Main] Starting application!")

	fmt.Println(coordinate.Coordinates)

	if err := api.NewApi().StartServer(); err != nil {
		panic("Panic starging server!")
	}

}
