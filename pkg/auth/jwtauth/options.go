package jwtauth

import (
	"time"
)

type Option struct {
	JwtSecretKey           string
	JwtAccessTokenExpires  time.Duration
	JwtRefreshTokenExpires time.Duration
	JwtIssuer              string
	JwtSubject             string
	JwtAudience            []string
}

// NewOption returns default jwt option.
func NewOption() *Option {
	return &Option{
		JwtAccessTokenExpires:  2 * time.Hour,
		JwtRefreshTokenExpires: 30 * 24 * time.Hour,
	}
}
