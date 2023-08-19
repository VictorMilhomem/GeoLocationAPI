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
	router.HandleFunc("/api/v1/store", controllers.GetAllStores).Methods("GET")
	router.HandleFunc("/api/v1/store", controllers.CreateStore).Methods("POST")
	router.HandleFunc("/api/v1/store/{id}", controllers.DeleteStore).Methods("DELETE")
	router.HandleFunc("/api/v1/store/{id}", controllers.UpdateStore).Methods("PUT")
	router.HandleFunc("/api/v1/store/{id}", controllers.GetStore).Methods("GET")
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
