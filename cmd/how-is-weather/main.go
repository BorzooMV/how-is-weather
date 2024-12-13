package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/BorzooMV/how-is-weather/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("couldn't load .env file:\n%v\n", err.Error())
	}

	appRouter := router.Router{}

	http.Handle("/api/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/weather") {
			appRouter.WeatherRouter(w, r)
			return
		}
		http.NotFound(w, r)
	}))

	fmt.Printf("Listening on port %v...\n", os.Getenv("SERVER_LISTENING_PORT"))
	err = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("SERVER_LISTENING_PORT")), nil)
	if err != nil {
		log.Fatalf("server couldn't start:\n%v\n", err.Error())
	}
}
