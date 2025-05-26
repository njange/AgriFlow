package main

import (
    "fmt"
    "net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Ingestor Service is healthy!")
}

func main() {
    http.HandleFunc("/health", healthCheck)
    fmt.Println("Ingestor Service running on port 8081")
    http.ListenAndServe(":8081", nil)
}