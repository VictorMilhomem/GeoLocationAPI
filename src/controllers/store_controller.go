package controllers

import (
	"encoding/json"
	"log"
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

func GetAllStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var store []models.Store

	database.DB.Find(&store)

	json.NewEncoder(w).Encode(store)
}

func getStoresLocation() ([]models.Point, error) {
	var stores []models.Store
	var points []models.Point

	database.DB.Find(&stores)
	for _, store := range stores {
		point, err := models.ConvertJSONToPoints(store.Addrs)
		if err != nil {
			return nil, err
		}
		points = append(points, *point)
	}
	return points, nil
}

func GetUserClosestStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user models.Point
	var store models.Store

	json.NewDecoder(r.Body).Decode(&user)
	// calculate the closest store
	store_locations, err := getStoresLocation()
	if err != nil {
		log.Println(err)
	}

	closestpoint := models.FindClosestPoint(user, store_locations)
	// find the store with the point
	addrs_json, err := models.ConvertPointtoJSONB(closestpoint)
	if err != nil {
		log.Println(err)
	}

	database.DB.Where("addrs = ?", addrs_json).Find(&store)
	json.NewEncoder(w).Encode(store)
}

func CreateStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var store models.Store

	json.NewDecoder(r.Body).Decode(&store)

	database.DB.Create(&store)

	json.NewEncoder(w).Encode(store)
}

func DeleteStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var store models.Store

	database.DB.Delete(&store, id)

	json.NewEncoder(w).Encode(store)
}

func UpdateStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)

	id := vars["id"]
	var store models.Store

	database.DB.First(&store, id)

	json.NewDecoder(r.Body).Decode(&store)

	database.DB.Save(&store)

	json.NewEncoder(w).Encode(store)
}
