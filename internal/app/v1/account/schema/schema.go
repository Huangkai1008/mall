package schema

import (
	"mall/internal/pkg/schema"
)

// CreateSchema is the account create schema.
type CreateSchema struct {
	schema.BaseSchema
	Username string `json:"username" binding:"required,max=127"`
	Email    string `json:"email" binding:"required,email,max=127"`
	Password string `json:"password" binding:"required,max=64"`
}

// LoginSchema is the account login schema.
type LoginSchema struct {
	schema.BaseSchema
	Username string `json:"username" binding:"required,max=127"`
	Password string `json:"password" binding:"required,max=64"`
}

type TokenSchema struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
