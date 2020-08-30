package validators

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"mall/internal/pkg/constant"
)

// Options for the validator.
type Options struct {
	// The locale to use for validation messages.
	Locale string
}

// NewOptions creates a new set of o for the HTTP server.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.Wrap(err, constant.LoadConfigError)
	}
	return o, err
}
