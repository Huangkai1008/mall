package user

import (
	"mall/internal/pkg/schema"
)

// RegisterSchema is the user register schema.
type RegisterSchema struct {
	schema.BaseSchema
	Username string `json:"username" validate:"required;max=127"`
	Email    string `json:"email" validate:"required;email;max=127"`
	Password string `json:"password" validate:"required;max=64"`
}

// Schema is the user return schema.
type Schema struct {
	schema.BaseSchema
	Username string `json:"username"`
	Email    string `json:"email"`
}
