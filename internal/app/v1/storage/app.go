package storage

import (
	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"

	"github.com/Huangkai1008/micro-kit/pkg/application"
	"github.com/Huangkai1008/micro-kit/pkg/transport/http"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

// New returns a new storage application.
func New(
	c *config.Config,
	logger *zap.Logger,
	httpServer *http.Server,
	minioCli *minio.Client,
) (*application.Application, error) {
	return application.New(
		c.App.Name,
		c.App.Version,
		logger,
		application.WithHttpServer(httpServer),
		application.WithMinioCli(minioCli),
	)
}

var ProviderSet = wire.NewSet(New)
