package api

import (
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/health", HealthCheck).Methods("GET")
    return router
}
