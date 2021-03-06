package encrypt

import "golang.org/x/crypto/bcrypt"

// GeneratePasswordHash generate password hash from raw password.
func GeneratePasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

// ComparePasswordHash compare password hash with raw password.
func ComparePasswordHash(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
