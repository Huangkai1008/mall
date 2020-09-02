package jwtauth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"

	"mall/pkg/auth"
)

var DefaultSigningMethod = jwt.SigningMethodHS256

type JwtAuth struct {
	*Options
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

// CreateRefreshToken create a new refresh token.
// The identity of this token, which can be any data that is json serializable.
func (j *JwtAuth) CreateRefreshToken(identity interface{}) (string, error) {
	claims := Claims{
		Identity:  identity,
		Fresh:     false,
		TokenType: RefreshTokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.JwtIssuer,
			Subject:   j.JwtSubject,
			Audience:  j.JwtAudience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.JwtRefreshTokenExpires)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(DefaultSigningMethod, claims)
	tokenString, err := token.SignedString(j.JwtSecretKey)
	return tokenString, err
}

// ParseToken parse token string to Claims.
func (j *JwtAuth) ParseToken(tokenString string) (auth.Claims, error) {
	return j.ParseJwtToken(tokenString)
}

func (j *JwtAuth) ParseJwtToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtSecretKey, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

var ProviderSet = wire.NewSet(wire.Bind(new(auth.Auth), new(*JwtAuth)), New)
