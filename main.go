package main

import (
    "log"
    "net/http"

    "go-messenger-api/api"
    "go-messenger-api/config"
)

func main() {
    config := config.LoadConfig()
    router := api.NewRouter()

    log.Printf("Starting server on %s\n", config.ServerAddress)
    log.Fatal(http.ListenAndServe(config.ServerAddress, router))
}