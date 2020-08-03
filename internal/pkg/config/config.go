package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const (
	DebugMode   = "debug"
	TestMode    = "testing"
	ReleaseMode = "release"
)

// New returns a new viper config.
func New(path string) (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)

	// Get basic configs from toml file.
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	if err = v.ReadInConfig(); err == nil {
		fmt.Printf("Use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	// Get secure configs from dotenv file
	v.SetConfigFile(".env")
	v.AutomaticEnv()
	if err = v.MergeInConfig(); err != nil {
		return nil, err
	}

	v.WatchConfig()
	return v, err
}

var ProviderSet = wire.NewSet(New)
