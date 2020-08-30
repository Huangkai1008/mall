package validators

import (
	"net/http"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validate *validator.Validate
	trans    ut.Translator
	locale   string
}

func New(o *Options) (*CustomValidator, error) {
	validate := validator.New()
	trans, err := registerTranslation(validate, o.Locale)
	if err != nil {
		return nil, err
	}
	return &CustomValidator{
		validate: validate,
		locale:   o.Locale,
		trans:    trans,
	}, nil
}

func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.validate.Struct(i); err != nil {
		var builder strings.Builder
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			msg := err.Translate(v.trans)
			builder.WriteString(msg)
			builder.WriteString(", ")
		}
		return echo.NewHTTPError(http.StatusUnprocessableEntity, builder.String())
	}
	return nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
