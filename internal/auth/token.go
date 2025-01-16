package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
)

func (auth *Auth) createToken(claims jwt.Claims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(auth.secretKey))
	return token, err
}

func (auth *Auth) CreateWsAccessToken(userID string) (string, error) {
	claims := newClaims("", userID, "", "", auth.wsAccessTokenDuration)
	return auth.createToken(claims)
}

func (auth *Auth) CreateAccessToken(sessionID, userID, firstName, lastName string) (string, error) {
	claims := newClaims(sessionID, userID, firstName, lastName, auth.accessTokenDuration)
	return auth.createToken(claims)
}

func (auth *Auth) CreateRefreshToken(sessionID, userID, firstName, lastName string) (string, error) {
	claims := newClaims(sessionID, userID, firstName, lastName, auth.refreshTokenDuration)
	return auth.createToken(claims)
}

func (auth *Auth) VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return &Claims{}, ErrInvalidToken
	}

	return claims, nil
}
