package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AppToken is a struct that contains the JWT token
type AppToken struct {
	Token          string    `json:"token"`
	TokenType      string    `json:"type"`
	ExpirationTime time.Time `json:"expitationTime"`
}

// Claims is a struct that contains the claims of the JWT
type Claims struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
