package user

import (
	"mall/internal/pkg/schema"
)

// RegisterSchema is the user register schema.
type RegisterSchema struct {
	schema.BaseSchema
	Username string `json:"username" binding:"required,max=127"`
	Email    string `json:"email" binding:"required,email,max=127"`
	Password string `json:"password" binding:"required,max=64"`
}

// Schema is the user return schema.
type Schema struct {
	schema.BaseSchema
	Username string `json:"username"`
	Email    string `json:"email"`
}
