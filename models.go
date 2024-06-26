package autoderm_go_client

import "github.com/google/uuid"

type QueryResponse struct {
	Success     bool         `json:"success"`
	Message     string       `json:"message"`
	ID          string       `json:"id"` // uuid.UUID on success or empty string on failure
	Predictions []Prediction `json:"predictions"`
}

type Prediction struct {
	Confidence       float64   `json:"confidence"`
	Icd              string    `json:"icd"`
	Name             string    `json:"name"`
	ClassificationID uuid.UUID `json:"classificationId"`
	ReadMoreURL      string    `json:"readMoreUrl"`
}
