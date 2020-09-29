package logging

import (
	"go.uber.org/zap"

	"github.com/Huangkai1008/micro-kit/pkg/logging"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

// NewLogger creates a new logger.
func NewLogger(c *config.Config) (*zap.Logger, error) {
	return logging.New(
		c.Log.FileName,
		logging.WithLevel(c.Log.Level),
		logging.WithStdout(c.Log.Stdout),
	)
}
