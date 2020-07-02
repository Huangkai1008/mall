package schema

import (
	"strings"

	"github.com/go-playground/validator/v10"

	"mall/internal/pkg/ecode"
	"mall/internal/pkg/util/field"
	validate "mall/internal/pkg/validator"
)

type Schema interface {
	Validate(errs validator.ValidationErrors) ecode.MallError
}

type BaseSchema struct {
}

func (r *BaseSchema) Validate(errs validator.ValidationErrors) ecode.MallError {
	var (
		mallError ecode.MallError
		builder   strings.Builder
	)

	for _, err := range errs {
		msg := err.Translate(validate.Trans)
		builder.WriteString(msg)
		builder.WriteString(", ")
	}
	mallError.Message = builder.String()
	return mallError
}

type ReturnSchema struct {
	ID        uint           `json:"id"`
	CreatedAt field.JsonTime `json:"create_time"`
	UpdatedAt field.JsonTime `json:"update_time"`
}
