package api

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"

    "go-messenger-api/handlers"
)

func NewRouter() http.Handler {
    router := mux.NewRouter()

    router.HandleFunc("/api/hello", handlers.HelloHandler).Methods("GET")
    // Добавьте больше маршрутов здесь

    // Настройка CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5173"}, // Разрешите origin вашего фронтенда
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    })

    return c.Handler(router)
}