package jwtauth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var DefaultSigningMethod = jwt.SigningMethodHS256

type JwtAuth struct {
	*Option
}

// New returns new JwtAuth.
func New() *JwtAuth {
	return &JwtAuth{}
}

// CreateAccessToken create a new access token.
// The identity of this token, which can be any data that is json serializable.
func (j *JwtAuth) CreateAccessToken(identity interface{}, fresh bool) (string, error) {
	claims := Claims{
		Identity:  identity,
		Fresh:     fresh,
		TokenType: AccessTokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.JwtIssuer,
			Subject:   j.JwtSubject,
			Audience:  j.JwtAudience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.JwtAccessTokenExpires)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(DefaultSigningMethod, claims)
	tokenString, err := token.SignedString(j.JwtSecretKey)
	return tokenString, err
}
