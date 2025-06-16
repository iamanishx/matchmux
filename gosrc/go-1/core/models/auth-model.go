package models

import "github.com/google/uuid"


type Credentials struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Otpverification struct {
	Code      string `json:"code"`
	UserID    uuid.UUID `json:"user_id"`
}