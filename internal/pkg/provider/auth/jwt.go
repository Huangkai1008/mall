package auth

import (
	"github.com/Huangkai1008/micro-kit/pkg/auth/jwtauth"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

// NewJwtAuth returns new JwtAuth.
func NewJwtAuth(c *config.Config) *jwtauth.JwtAuth {
	return jwtauth.New(
		c.Jwt.SecretKey,
		jwtauth.WithIssuer(c.Jwt.Issuer),
		jwtauth.WithAudience(c.Jwt.Audience...),
		jwtauth.WithSubject(c.Jwt.Subject),
	)
}
