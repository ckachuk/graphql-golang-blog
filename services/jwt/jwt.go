package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JwtCustomClaim struct {
	Username string 
	Exp 	 int64	
	jwt.RegisteredClaims
}
var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("sampleSecretKey")
	if secret == "" {
		return "aSecret"
	}
	return secret
}


func GenerateToken(username string) (string, error) {
	claims:= JwtCustomClaim{
		username,
		time.Now().Add(time.Hour * 24).Unix(),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err:= token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString,nil
}



func JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}