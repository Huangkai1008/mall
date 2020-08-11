package account

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"mall/internal/pkg/application"
	"mall/internal/pkg/constant"
	"mall/internal/pkg/transport/http"
)

type Options struct {
	Name   string
	Locale string
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, constant.LoadConfigError)
	}

	logger.Info(constant.AppConfigOk)

	return o, err
}

// New returns a new account application.
func New(
	o *Options,
	logger *zap.Logger,
	httpServer *http.Server,
) (*application.Application, error) {
	return application.New(
		o.Name,
		o.Locale,
		logger,
		application.WithHttpServer(httpServer),
	)
}

var Tables = []interface{}{&Account{}}
var ProviderSet = wire.NewSet(New, NewOptions, wire.NewSet(wire.Value(Tables)))
