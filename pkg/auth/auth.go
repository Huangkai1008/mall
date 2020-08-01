package auth

type Auth interface {
	// CreateToken create a new token.
	// The identity of this token, which can be any data that is json serializable.
	CreateToken(identity interface{}, fresh bool) (string, error)

	ParseToken(token string)
}
