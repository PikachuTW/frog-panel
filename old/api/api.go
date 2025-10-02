package api

import (
	"github.com/go-fuego/fuego"
)

func New() *fuego.Server {
	s := fuego.NewServer()

	return s
}
