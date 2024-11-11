package model

// UserData model is
type UserData struct {
	Subject string   `json:"subject" validate:"required"`
	Email   []string `json:"email" validate:"required"`
	Message string   `json:"message" validate:"required"`
}
