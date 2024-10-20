package api

import (
    // "net/http"
    "github.com/gorilla/mux"
)

// NewRouter creates and returns a new router
func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Define your routes and handlers
    router.HandleFunc("/users", GetUsersHandler).Methods("GET")

    return router
}
