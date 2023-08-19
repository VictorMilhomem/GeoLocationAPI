package server

import (
	"log"
	"net/http"

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
	log.Printf("Running server on port: %s", s.listenAddr)
	routes.HandleRequest()
	log.Fatal(http.ListenAndServe(s.listenAddr, nil))
}
