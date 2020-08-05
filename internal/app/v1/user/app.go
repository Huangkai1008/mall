package user

import (
	"mall/internal/pkg/application"
	"mall/internal/pkg/constant"
	"mall/internal/pkg/transport/http"

	"github.com/google/wire"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Options struct {
	Name string
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

// New returns a new user application.
func New(
	o *Options,
	logger *zap.Logger,
	httpServer *http.Server,
	minioCli *minio.Client,
) (*application.Application, error) {
	return application.New(
		o.Name,
		logger,
		application.WithHttpServer(httpServer),
		application.WithMinioCli(minioCli),
	)
}

var ProviderSet = wire.NewSet(New, NewOptions)