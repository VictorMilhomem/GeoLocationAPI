package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VictorMilhomem/GeoLocationApi/src/database"
	"github.com/VictorMilhomem/GeoLocationApi/src/models"
	"github.com/gorilla/mux"
)

func GetStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)

	id := vars["id"]

	var store models.Store

	database.DB.First(&store, id)

	json.NewEncoder(w).Encode(store)
}
