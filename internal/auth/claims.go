package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID    string `json:"userId,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	SessionID string `json:"sessionId,omitempty"`
	jwt.RegisteredClaims
}

func newClaims(sessionID, userID, firstName, lastName string, duration time.Duration) Claims {
	claims := Claims{
		UserID:    userID,
		FirstName: firstName,
		LastName:  lastName,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return claims
}
