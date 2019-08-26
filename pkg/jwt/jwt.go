package jwt

import (
	"fmt"
	"time"

	"boilerplate/pkg/logger"

	"github.com/dgrijalva/jwt-go"
)

// UserJWTClaims is the structure of a JWT for a User
type UserJWTClaims struct {
	UserID int `json:"uid"`
	jwt.StandardClaims
}

// GeneratedResponse is the response of a generated token, usually used in http response
type GeneratedResponse struct {
	Expires int64  `json:"expires"`
	Token   string `json:"token"`
}

// Generate will create a JWT
func Generate(userID int) (GeneratedResponse, error) {
	var response GeneratedResponse

	key, _ := GetPrivateKey()
	expires := time.Now().AddDate(0, 0, 1) // 1 day

	// Create the Claims
	claims := UserJWTClaims{
		userID,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expires.Unix(),
			Issuer:    "api",
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	var err error
	token.Signature, err = token.SignedString(key)
	if err != nil {
		logger.Error("%+v", fmt.Errorf("Error signing token: %v", err))
		return response, err
	}

	response = GeneratedResponse{
		Expires: expires.Unix(),
		Token:   token.Signature,
	}

	return response, nil
}
