package user

import (
	"go.uber.org/zap"

	"mall/internal/pkg/util/encrypt"
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
	hashPassword, err := encrypt.GeneratePasswordHash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPassword
	err = s.repo.Create(user)
	return user, err
}
