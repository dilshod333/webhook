package main

import (
    "log"
    "net/http"

    "conn/internal/handlers"
    "conn/internal/mongo"
)

func main() {
   
    if err := mongo.InitMongoDB(); err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    defer mongo.DisconnectMongoDB()

    http.HandleFunc("/webhook", handlers.WebhookHandler)
    log.Println("Listening on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
