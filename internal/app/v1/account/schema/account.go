package schema

// AccountRegSchema is the account create schema.
type AccountRegSchema struct {
	Username string `json:"username" validate:"required,max=127"`
	Email    string `json:"email" validate:"required,email,max=127"`
	Password string `json:"password" validate:"required,max=64"`
}

// AccountLoginSchema is the account login schema.
type AccountLoginSchema struct {
	Username string `json:"username" validate:"required,max=127"`
	Password string `json:"password" validate:"required,max=64"`
}

type AccountTokenSchema struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
