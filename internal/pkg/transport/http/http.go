// Package http implements the HTTP transport protocol.
package http

import (
	"context"
	"github.com/google/wire"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// Server is the HTTP server.
type Server struct {
	httpServer *http.Server
	router     *gin.Engine
	logger     *zap.Logger
	o          *Options
}

// New creates a new HTTP server.
func New(o *Options, logger *zap.Logger, router *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           o.Addr(),
			ReadTimeout:    o.ReadTimeout * time.Second,
			WriteTimeout:   o.WriteTimeout * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		router: router,
		logger: logger.With(zap.String("type", "http.Server")),
		o:      o,
	}
}

// Start http server.
func (s *Server) Start() error {
	s.httpServer.Handler = s.router
	s.logger.Info("HTTP server starting ...", zap.String("addr", s.o.Addr()))

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("Start HTTP server error", zap.Error(err))
		}
	}()
	return nil
}

// Stop http server.
func (s *Server) Stop() error {
	s.logger.Info("HTTP server stopping ...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Shutdown HTTP server error")
	}
	return nil
}

var ProviderSet = wire.NewSet(New, NewOptions, NewRouter)
