package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Credentials struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Otpverification struct {
	Code   string    `json:"code"`
	UserID uuid.UUID `json:"user_id"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}
