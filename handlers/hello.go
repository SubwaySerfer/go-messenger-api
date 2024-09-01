package handlers

import (
    "encoding/json"
    "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"message": "Hello from Go API!"}
    json.NewEncoder(w).Encode(response)
}