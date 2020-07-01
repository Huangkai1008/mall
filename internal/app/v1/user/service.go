package user

import (
	"go.uber.org/zap"

	"golang.org/x/crypto/bcrypt"

	repo "mall/internal/pkg/repository"
	srv "mall/internal/pkg/service"
)

type Service struct {
	logger *zap.Logger
	srv.CRUDService
}

func NewService(logger *zap.Logger, repo repo.GormRepository) *Service {
	return &Service{
		logger: logger.With(zap.String("type", "UserService")),
		CRUDService: srv.CRUDService{
			Repo: &repo,
		},
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

	err = s.Repo.Create(user)
	return user, err
}
