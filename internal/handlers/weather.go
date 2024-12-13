package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/BorzooMV/how-is-weather/internal/services"
)

type CurrentCondition struct {
	DatetimeEpoch int64   `json:"datetimeEpoch"`
	Temperature   float64 `json:"temp"`
	FeelsLike     float64 `json:"feelslike"`
	Humidity      float64 `json:"humidity"`
	Conditions    string  `json:"conditions"`
}
type Response struct {
	Address          string           `json:"address"`
	Timezone         string           `json:"timezone"`
	Description      string           `json:"description"`
	CurrentCondition CurrentCondition `json:"currentConditions"`
}

func GetWeather(w http.ResponseWriter, r *http.Request, city string) {
	var response Response

	ctx := context.Background()
	redisClient := services.ConnectRedis()
	defer redisClient.Close()

	cachedValue, err := redisClient.Get(ctx, city).Result()
	// Fetch from API then cache the response
	if err != nil {
		err := getData(city, w, &response)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		data, err := json.Marshal(response)
		if err != nil {
			http.Error(w, fmt.Sprintf("can't marshal json:\n%v\n", err.Error()), http.StatusInternalServerError)
			return
		}

		_, err = redisClient.Set(ctx, city, data, 0).Result()
		if err != nil {
			fmt.Printf("Failed to cache the data:\n%v\n", err.Error())
		}
	} else {
		// Server from the cache
		err := json.Unmarshal([]byte(cachedValue), &response)
		if err != nil {
			http.Error(w, fmt.Sprintf("couldn't unmarshal json:\n%v\n", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getData(city string, w http.ResponseWriter, response *Response) error {
	resp, err := http.Get(fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%v?unitGroup=metric&include=current&key=%v&contentType=json", city, os.Getenv("VISUAL_CROSSING_API_KEY")))
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't get weather data:\n%v\n", err.Error()), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("not found")
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't read requested body: \n%v\n", err.Error()), http.StatusInternalServerError)
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't unmarshal json:\n%v\n", err.Error()), http.StatusInternalServerError)
	}

	return nil
}
