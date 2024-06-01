package model

import (
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}
