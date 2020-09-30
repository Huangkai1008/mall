package config

import (
	"fmt"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/Huangkai1008/micro-kit/pkg/auth/jwtauth"
	"github.com/Huangkai1008/micro-kit/pkg/database/gorm"
	"github.com/Huangkai1008/micro-kit/pkg/logging"
	"github.com/Huangkai1008/micro-kit/pkg/registry/consul"
	"github.com/Huangkai1008/micro-kit/pkg/storage/minio"
	"github.com/Huangkai1008/micro-kit/pkg/transport/http"

	"github.com/Huangkai1008/mall/internal/pkg/constant"
)

// Config is the config of application.
type Config struct {
	App      App
	Log      Log
	Database Database
	HTTP     HTTP
	Minio    Minio
	Jwt      jwtauth.Options
	Consul   consul.Options
}

type App struct {
	Name    string
	Version string
	Locale  string
}

type Log struct {
	logging.Options `mapstructure:",squash"`
}

type Database struct {
	gorm.Options `mapstructure:",squash"`
}

type HTTP struct {
	http.Options `mapstructure:",squash"`
}

type Minio struct {
	minio.Options `mapstructure:",squash"`
}

// New returns a new viper config.
func New(path string) (*Config, error) {
	var (
		err    error
		v      = viper.New()
		config *Config
	)

	// Get basic configs from toml file.
	v.SetConfigName(path)
	v.SetConfigType("toml")
	v.AddConfigPath(fmt.Sprintf("configs/%s", path))
	if err = v.ReadInConfig(); err == nil {
		fmt.Printf("Use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, errors.Wrap(err, constant.LoadConfigError)
	}

	// Get secure configs from dotenv file
	v.SetConfigFile(fmt.Sprintf("configs/%s/.env", path))
	v.AutomaticEnv()
	if err = v.MergeInConfig(); err != nil {
		return nil, errors.Wrap(err, constant.LoadConfigError)
	}
	for _, key := range v.AllKeys() {
		val := v.Get(key)
		v.Set(key, val)
	}
	v.WatchConfig()
	if err = v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, err
}

var ProviderSet = wire.NewSet(New)
