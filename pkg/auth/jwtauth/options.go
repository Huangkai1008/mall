package jwtauth

import (
	"time"

	"mall/internal/pkg/constant"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Options is a struct for specifying configuration options for the JWT.
type Options struct {
	JwtSecretKey           string
	JwtAccessTokenExpires  time.Duration
	JwtRefreshTokenExpires time.Duration
	JwtIssuer              string
	JwtSubject             string
	JwtAudience            []string
}

// NewOptions returns a new instance of the options with the given parameters.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	v.SetDefault("JwtAccessTokenExpires", 2*time.Hour)
	v.SetDefault("JwtRefreshTokenExpires", 30*24*time.Hour)

	if err = v.UnmarshalKey("jwt", o); err != nil {
		return nil, errors.Wrap(err, constant.LoadConfigError)
	}
	return o, err
}
