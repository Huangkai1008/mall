package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

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
	Locale  string
	Log
	HTTP
	Gorm
	Database
	Minio
}

// New returns new Config instance.
func New() (*Config, error) {
	var (
		err    error
		config *Config
	)

	v := viper.New()
	// Get basic configs from toml file
	v.AddConfigPath("configs")
	v.SetConfigType("toml")
	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}
	// Get secure configs from dotenv file
	v.SetConfigFile(".env")
	v.AutomaticEnv()
	if err = v.MergeInConfig(); err != nil {
		return nil, err
	}

	if err = v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return config, err
}

type Log struct {
	Level    int    // logging level
	FileName string // logging filename
}

type HTTP struct {
	HttpHost     string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Addr return the TCP address in the form "host:port"
func (h HTTP) Addr() string {
	return fmt.Sprintf("%s:%d", h.HttpHost, h.HttpPort)
}

type Gorm struct {
	MaxIdleConnections int
	MaxOpenConnections int
	EnableAutoMigrate  bool
}

type Database struct {
	User       string
	Password   string
	Host       string
	Port       int
	DBName     string
	Parameters string
}

func (d Database) DSN() string {
	const dsn = "%s:%s@tcp(%s:%d)/%s?%s"
	return fmt.Sprintf(dsn, d.User, d.Password, d.Host, d.Port, d.DBName, d.Parameters)
}

type Minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	Region          string
}
