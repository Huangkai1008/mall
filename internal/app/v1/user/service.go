package user

import (
	"go.uber.org/zap"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger *zap.Logger
	repo   Repository
}

func NewService(logger *zap.Logger, repo Repository) *Service {
	return &Service{
		logger: logger.With(zap.String("type", "UserService")),
		repo:   repo,
	}
}

// Create User.
func (s *Service) Create(user *User) (*User, error) {
	// Generate hash password to encrypt.
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashPassword)

	err = s.repo.Create(user)
	return user, err
}
