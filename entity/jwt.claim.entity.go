package entity

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	jwt.StandardClaims
	Id uint `json:"id"`
}