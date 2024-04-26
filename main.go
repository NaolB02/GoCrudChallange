package main

import (
    "log"
    "net/http"

    handlers "main/handlers"
    "github.com/rs/cors"
)

func main() {
    // Initialize the router
    router := handlers.NewRouter()

    // CORS middleware
    c := cors.AllowAll()

    // Handler with CORS support
    handler := c.Handler(router)

    // Start the server
    log.Fatal(http.ListenAndServe(":8080", handler))
}
