package api

import (
    "net/http"
    "fmt"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Server is running!")
}
