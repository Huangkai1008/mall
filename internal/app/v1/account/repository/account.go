package repository

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"

	repo "mall/internal/pkg/repository"
	metav1 "mall/pkg/meta/v1"
)

type AccountRepository interface {
	Find(conditions interface{}) (record metav1.Resource, err error)

	Exist(condition interface{}) (bool, error)

	Create(record metav1.Resource) error
}

type accountRepository struct {
	*repo.GormRepository
}

// NewAccountRepository returns new account AccountRepository.
func NewAccountRepository(logger *zap.Logger, db *gorm.DB) AccountRepository {
	return &accountRepository{
		repo.NewGormRepository(db, logger),
	}
}

var ProviderSet = wire.NewSet(NewAccountRepository)
