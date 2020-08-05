package service

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"mall/internal/app/v1/user"
	"mall/internal/app/v1/user/repository"
	"mall/internal/pkg/util/encrypt"
)

type Service struct {
	logger *zap.Logger
	repo   repository.Repository
}

func NewService(logger *zap.Logger, repo repository.Repository) *Service {
	return &Service{
		logger: logger.With(zap.String("type", "UserService")),
		repo:   repo,
	}
}

func (s *Service) Create(user *user.User) (*user.User, error) {
	// Generate hash password to encrypt.
	hashPassword, err := encrypt.GeneratePasswordHash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPassword
	err = s.repo.Create(user)
	return user, err
}

func (s *Service) validatePassword(user *user.User, password string) (bool, error) {
	hashPassword, err := encrypt.GeneratePasswordHash(password)
	if err != nil {
		return false, err
	}
	return user.Password == hashPassword, nil
}

var ProviderSet = wire.NewSet(NewService)
