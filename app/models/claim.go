package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claim : struct to validate jwt
type Claim struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	jwt.StandardClaims
}
