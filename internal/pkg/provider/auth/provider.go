package auth

import (
	"github.com/Huangkai1008/micro-kit/pkg/auth/jwtauth"
	"github.com/google/wire"

	"github.com/Huangkai1008/micro-kit/pkg/auth"
)

var ProviderSet = wire.NewSet(wire.Bind(new(auth.Auth), new(*jwtauth.JwtAuth)), NewJwtAuth)
