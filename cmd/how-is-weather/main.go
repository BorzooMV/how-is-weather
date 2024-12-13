package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("couldn't load .env file:\n%v\n", err.Error())
	}

	fmt.Printf("Listening on port %v...\n", os.Getenv("SERVER_LISTENING_PORT"))
	err = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("SERVER_LISTENING_PORT")), nil)
	if err != nil {
		log.Fatalf("server couldn't start:\n%v\n", err.Error())
	}
}
