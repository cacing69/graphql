package lib

import (
	"github.com/cacing69/graphql/entity"
	"github.com/cacing69/graphql/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret")
var APPLICATION_NAME = "api-source"

func ClaimToken(user model.User) (string, error) {
	claims := entity.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Id: user.Id,
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	return token.SignedString(JWT_SIGNATURE_KEY)
}


