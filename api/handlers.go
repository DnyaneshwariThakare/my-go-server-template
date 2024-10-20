package api

import (
    "encoding/json"
    "log"
    "my-go-server/internal/repository"
    "net/http"
)

// GetUsersHandler is an HTTP handler to retrieve users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    users, err := repository.GetUsers() // Capture users and error
    if err != nil {
        log.Printf("Error fetching users: %v", err)
        http.Error(w, "Error fetching users", http.StatusInternalServerError)
        return
    }

    // Log the users in a readable format
    log.Printf("Fetched users: %+v", users)

    // Convert the users to JSON and send the response
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(users); err != nil {
        log.Printf("Error encoding users to JSON: %v", err)
        http.Error(w, "Error encoding users to JSON", http.StatusInternalServerError)
    }

    // Optionally, if you want to log the JSON data before sending it to the client
    jsonData, err := json.Marshal(users)
    if err == nil {
        log.Printf("JSON data to be sent: %s", string(jsonData))
    }
}
