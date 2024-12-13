package handlers

import (
	"encoding/json"
	"net/http"
)

func GetWeather(w http.ResponseWriter, r *http.Request, city string) {
	var Response struct {
		City string `json:"city"`
	}

	Response.City = city

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response)
}
