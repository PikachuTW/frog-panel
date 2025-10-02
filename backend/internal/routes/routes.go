package routes

import (
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func SetupRoutes(s *fuego.Server) {
	SetupInfoHandlers(fuego.Group(s, "/info", option.Tags("Info")))
}
