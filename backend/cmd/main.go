package main

import (
	"frog-panel/internal/routes"

	"github.com/go-fuego/fuego"
)

func main() {
	s := fuego.NewServer()

	routes.SetupRoutes(s)

	s.Run()
}
