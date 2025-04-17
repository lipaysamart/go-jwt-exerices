package jtoken

import (
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lipaysamart/go-jwt-exerices/pkg/utils"
)

func GenerateAccessToken(payload map[string]interface{}) string {
	payload["type"] = "x-access"
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * 5 * 60 * 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		log.Printf("Error generating token: %v\n", err)
		return ""
	}

	return token
}

func RefreshToken(payload map[string]interface{}) string {
	payload["type"] = "x-refresh"
	tokenContent := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Second * 10 * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		log.Printf("Error refresh token: %v\n", err)
		return ""
	}

	return token
}

func ValidateToken(jwtToken string) (map[string]interface{}, error) {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	var data map[string]interface{}
	utils.Copy(&data, tokenData["payload"])
	return data, nil
}
