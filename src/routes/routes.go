package routes

import (
	"github.com/VictorMilhomem/GeoLocationApi/src/controllers"
	"github.com/VictorMilhomem/GeoLocationApi/src/middleware"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.Use(middleware.ContetTypeMiddleware)
	router.HandleFunc("/api/v1/store/{id}", controllers.GetStore)
}
