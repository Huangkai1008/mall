package service

import (
	"github.com/google/wire"
	"go.uber.org/zap"

	"mall/internal/app/v1/account"
	"mall/internal/app/v1/account/repository"
	"mall/internal/pkg/util/encrypt"
	"mall/pkg/auth"
	"mall/pkg/auth/jwtauth"
)

type AccountService struct {
	logger *zap.Logger
	repo   repository.AccountRepository
	auth   auth.Auth
}

func NewAccountService(logger *zap.Logger, repo repository.AccountRepository, auth auth.Auth) *AccountService {
	return &AccountService{
		logger: logger.With(zap.String("type", "UserService")),
		repo:   repo,
		auth:   auth,
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
//	a, err := s.repo.Find(map[string]string{"username": username})
//	if err != nil {
//		return nil, error.
//	}
//	if a == nil {
//		return nil, error.New(constant.AccountNotExist)
//	}
//
//	// Check password.
//	if err = encrypt.ComparePasswordHash(password, a.Password); err != nil {
//		return nil, error.New(constant.AccountNotCorrectPassword)
//	}
//
//	accessToken, err := s.auth.CreateAccessToken(a.ID, true)
//	if err != nil {
//		return nil, err
//	}
//	refreshToken, err := s.auth.CreateRefreshToken(a.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	return &schema.AccountTokenSchema{
//		AccessToken:  accessToken,
//		RefreshToken: refreshToken,
//	}, nil
//}

func (s *AccountService) RefreshToken() {

}

var ProviderSet = wire.NewSet(NewAccountService, wire.Bind(new(auth.Auth), new(*jwtauth.JwtAuth)), jwtauth.New)
