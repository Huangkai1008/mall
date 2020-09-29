package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/Huangkai1008/micro-kit/pkg/registry"
	"github.com/Huangkai1008/micro-kit/pkg/transport/http"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

// NewHTTPServer creates a new HTTP server.
func NewHTTPServer(
	c *config.Config,
	logger *zap.Logger,
	router *echo.Echo,
	registrar registry.Registrar) *http.Server {
	return http.New(
		logger,
		router,
		registrar,
		http.WithHost(c.HTTP.Host),
		http.WithPort(c.HTTP.Port),
		http.WithMode(c.HTTP.Mode),
	)
}
