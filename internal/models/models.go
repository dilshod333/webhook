package models

type RepositoryEvent struct {
    Action     string `json:"action"`
    Repository struct {
        Name     string `json:"name"`
        FullName string `json:"full_name"`
    } `json:"repository"`
}
