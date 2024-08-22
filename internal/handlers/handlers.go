package handlers

import (
    "encoding/json"
    "io"
    "log"
    "net/http"

    "conn/internal/models"
    "conn/internal/mongo"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
    var event models.RepositoryEvent

    if r.Method != http.MethodPost {
        http.Error(w, "wronggg request method", http.StatusMethodNotAllowed)
        return
    }

   
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("Error reading request body: %v", err)
        http.Error(w, "Unable to read request body", http.StatusBadRequest)
        return
    }

   
    if err := json.Unmarshal(body, &event); err != nil {
        log.Printf("Error parsing JSON: %v", err)
        http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
        return
    }

    log.Printf("Received repository event: %s for repository %s", event.Action, event.Repository.FullName)

  
    if err := mongo.SaveToMongoDB(event); err != nil {
        log.Printf("Error saving to MongoDB: %v", err)
        http.Error(w, "wrrronngg to save event to database", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
