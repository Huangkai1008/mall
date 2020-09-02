package config

import (
	"fmt"
	"time"

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

	// Set viper defaults.
	setDefaultValues(v)

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
	return v, err
}

func setDefaultValues(v *viper.Viper) {
	v.SetDefault("jwt.access_token_expires", 2*time.Hour)
	v.SetDefault("jwt.refresh_token_expires", 30*24*time.Hour)

	v.SetDefault("consul.enable_health_check", true)
	v.SetDefault("consul.health_check_interval", 10)
	v.SetDefault("consul.deregister_critical_service_after", 60)
	v.SetDefault("consul.heart_beat", true)
}

var ProviderSet = wire.NewSet(New)
