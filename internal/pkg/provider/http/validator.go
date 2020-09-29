package http

import (
	"github.com/Huangkai1008/micro-kit/pkg/transport/http/validator"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

func NewValidator(c *config.Config) (*validator.CustomValidator, error) {
	return validator.New(c.App.Locale)
}
