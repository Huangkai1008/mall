package schema

import (
	"mall/internal/pkg/schema"
)

// AccountRegisterSchema is the account create schema.
type AccountRegisterSchema struct {
	schema.BaseSchema
	Username string `json:"username" binding:"required,max=127"`
	Email    string `json:"email" binding:"required,email,max=127"`
	Password string `json:"password" binding:"required,max=64"`
}

// AccountLoginSchema is the account login schema.
type AccountLoginSchema struct {
	schema.BaseSchema
	Username string `json:"username" binding:"required,max=127"`
	Password string `json:"password" binding:"required,max=64"`
}

type AccountTokenSchema struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
