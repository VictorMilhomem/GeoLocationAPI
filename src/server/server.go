package server

import (
	"log"

	"github.com/VictorMilhomem/GeoLocationApi/src/database"
	"github.com/VictorMilhomem/GeoLocationApi/src/routes"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	log.Println("Connecting to database...")
	database.ConnectDB()

	log.Printf("Running server on port %s", s.listenAddr)
	routes.HandleRequest(s.listenAddr)
}
