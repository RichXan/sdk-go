package dto

import (
	"github.com/golang-jwt/jwt/v4"
)

type AccessTokenClaims struct {
	jwt.StandardClaims
}
