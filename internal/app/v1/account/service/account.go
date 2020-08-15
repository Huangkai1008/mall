package service

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"mall/internal/app/v1/account"
	"mall/internal/app/v1/account/repository"
	"mall/internal/pkg/util/encrypt"
)

type AccountService struct {
	logger *zap.Logger
	repo   repository.AccountRepository
}

func NewAccountService(logger *zap.Logger, repo repository.AccountRepository) *AccountService {
	return &AccountService{
		logger: logger.With(zap.String("type", "UserService")),
		repo:   repo,
	}
}

func (s *AccountService) Create(account *account.Account) (*account.Account, error) {
	// Generate hash password to encrypt.
	hashPassword, err := encrypt.GeneratePasswordHash(account.Password)
	if err != nil {
		return nil, err
	}

	account.Password = hashPassword
	err = s.repo.Create(account)
	return account, err
}

//func (s *AccountService) Login(username, password string) (*schema.AccountTokenSchema, error) {
//	account, err := s.repo.Find(map[string]string{"username": username})
//	if err != nil {
//		return nil, err
//	}
//
//	// Check password.
//	if err := encrypt.CheckPasswordHash(password, account.Password); err != nil {
//		return nil, err
//	}
//
//	return account, nil
//}

func (s *AccountService) RefreshToken() {

}

func (s *AccountService) validatePassword(user *account.Account, password string) (bool, error) {
	hashPassword, err := encrypt.GeneratePasswordHash(password)
	if err != nil {
		return false, err
	}
	return user.Password == hashPassword, nil
}

var ProviderSet = wire.NewSet(NewAccountService)
