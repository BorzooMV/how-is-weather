package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetWeather(w http.ResponseWriter, r *http.Request, city string) {
	type CurrentCondition struct {
		DatetimeEpoch int64   `json:"datetimeEpoch"`
		Temperature   float64 `json:"temp"`
		FeelsLike     float64 `json:"feelslike"`
		Humidity      float64 `json:"humidity"`
		Conditions    string  `json:"conditions"`
	}
	var Response struct {
		Address          string           `json:"address"`
		Timezone         string           `json:"timezone"`
		Description      string           `json:"description"`
		CurrentCondition CurrentCondition `json:"currentConditions"`
	}

	resp, err := http.Get(fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%v?unitGroup=metric&include=current&key=%v&contentType=json", city, os.Getenv("VISUAL_CROSSING_API_KEY")))
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't get weather data:\n%v\n", err.Error()), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("couldn't get weather data, status:\n%v\t%v\n", resp.StatusCode, resp.Status), http.StatusInternalServerError)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't read requested body: \n%v\n", err.Error()), http.StatusInternalServerError)
	}

	err = json.Unmarshal(respBody, &Response)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't unmarshal json:\n%v\n", err.Error()), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response)
}
