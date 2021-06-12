package main

import (
	"log"
	"net/http"

	"github.com/cass-dlcm/SplatStatsGo/api_code"
)

func main() {
	log.Printf("Server started")

	router := api_code.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
