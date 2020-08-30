package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"mall/internal/pkg/router"
	"mall/internal/pkg/validators"
)

// NewRouter returns a new Echo router.
func NewRouter(o *Options, logger *zap.Logger, group router.Group, validator *validators.CustomValidator) (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Validator = validator
	apiGroup := e.Group("/api")
	v1Group := apiGroup.Group("/v1")
	{
		group(v1Group)
	}
	return e, nil
}
