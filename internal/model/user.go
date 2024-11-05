package model

// UserData model is
type UserData struct {
	FullName string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Message  string `json:"message" validate:"required"`
}
