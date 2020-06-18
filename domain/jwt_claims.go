package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
	jwt.StandardClaims
}

type ctxKey string

const JwtCtxKey ctxKey = "jwtClaims"
