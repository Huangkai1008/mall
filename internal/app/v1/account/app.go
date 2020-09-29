package account

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/Huangkai1008/micro-kit/pkg/application"
	"github.com/Huangkai1008/micro-kit/pkg/transport/http"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

// New returns a new account application.
func New(
	c *config.Config,
	logger *zap.Logger,
	httpServer *http.Server,
) (*application.Application, error) {
	return application.New(
		c.App.Name,
		c.App.Version,
		logger,
		application.WithHttpServer(httpServer),
	)
}

var Tables = []interface{}{&Account{}}
var ProviderSet = wire.NewSet(New, wire.NewSet(wire.Value(Tables)))
