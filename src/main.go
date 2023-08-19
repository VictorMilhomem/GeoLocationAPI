package main

import (
	"github.com/VictorMilhomem/GeoLocationApi/src/server"
)

func main() {
	server := server.NewAPIServer(":8000")
	server.Run()
}
