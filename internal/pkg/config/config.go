package config

import "github.com/spf13/viper"

type RunMode string

const (
	DebugMode   = "debug"
	TestMode    = "testing"
	ReleaseMode = "release"
)

// Config contains all the configs in the application.
type Config struct {
	Version string
	RunMode string
}

// New returns new Config instance.
func New() (*Config, error) {
	var (
		err    error
		config *Config
	)

	v := viper.New()
	v.SetConfigFile(".env")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err = v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, err
}
