package routes

import (
	"log"
	"net/http"

	"github.com/VictorMilhomem/GeoLocationApi/src/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest(listenAddr string) {
	router := mux.NewRouter()
	// router.Use(middleware.ContetTypeMiddleware)
	router.HandleFunc("/api/v1/store/{id}", controllers.GetStore).Methods("GET")
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
