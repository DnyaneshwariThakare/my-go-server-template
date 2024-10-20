// cmd/server/main.go
package main

import (
    "log"
    "my-go-server/api"
    "my-go-server/config"
    "net/http"
)

func main() {
    // Initialize the database connection
    config.InitDB()

    // Initialize the router
    router := api.NewRouter()

    // Start the server
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
