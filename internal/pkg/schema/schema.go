package schema

import (
	"github.com/go-playground/validator/v10"

	"mall/internal/pkg/ecode"
	"mall/internal/pkg/util/field"
)

type Schema interface {
	Validate(errs validator.ValidationErrors) ecode.MallError
}

type BaseSchema struct {
}

func (r *BaseSchema) Validate(errs validator.ValidationErrors) ecode.MallError {
	var mallError ecode.MallError
	mallError.Message = errs.Error()
	return mallError
}

type ReturnSchema struct {
	ID        uint           `json:"id"`
	CreatedAt field.JsonTime `json:"create_time"`
	UpdatedAt field.JsonTime `json:"update_time"`
}
