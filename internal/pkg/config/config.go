package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"mall/internal/pkg/constant"
)

// New returns a new viper config.
func New(path string) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
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

	v.WatchConfig()
	return v, err
}

var ProviderSet = wire.NewSet(New)
