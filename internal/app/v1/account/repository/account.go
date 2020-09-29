package repository

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"

	metav1 "github.com/Huangkai1008/micro-kit/pkg/meta/v1"
	repo "github.com/Huangkai1008/micro-kit/pkg/repository"

	"github.com/Huangkai1008/mall/internal/app/v1/account"
)

type AccountRepository interface {
	Find(conditions interface{}) (account *account.Account, err error)

	Exist(condition interface{}) (bool, error)

	Create(record metav1.Resource) error
}

type accountRepository struct {
	*repo.GormRepository
}

func (r *accountRepository) Find(conditions interface{}) (account *account.Account, err error) {
	err = r.Db.Where(conditions).Limit(1).Find(&account).Error
	return
}

// NewAccountRepository returns new account AccountRepository.
func NewAccountRepository(logger *zap.Logger, db *gorm.DB) AccountRepository {
	return &accountRepository{
		repo.NewGormRepository(db, logger),
	}
}

var ProviderSet = wire.NewSet(NewAccountRepository)
