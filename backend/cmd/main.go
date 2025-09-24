package main

import (
	"frog-panel/internal/api"
)

type Version struct {
	Stable  bool   `json:"stable"`
	Version string `json:"version"`
}

func main() {
	router := api.New()

	err := router.Run()

	if err != nil {
		panic(err)
	}

	// client := client.New()
	// var data []Version
	// client.GetJSON("https://meta.fabricmc.net/v2/versions/game", &data)
	// data = slices.DeleteFunc(data, func(version Version) bool {
	// 	return !version.Stable
	// })
	// fmt.Println(data)
}
