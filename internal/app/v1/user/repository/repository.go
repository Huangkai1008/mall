package repository

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"

	repo "mall/internal/pkg/repository"
	metav1 "mall/pkg/meta/v1"
)

type Repository interface {
	GetOne(conditions interface{}) (record metav1.Resource, err error)

	Exist(condition interface{}) (bool, error)

	Create(record metav1.Resource) error
}

type repository struct {
	*repo.GormRepository
}

// NewRepository returns new user Repository.
func NewRepository(logger *zap.Logger, db *gorm.DB) Repository {
	return &repository{
		repo.NewGormRepository(db, logger),
	}
}

var ProviderSet = wire.NewSet(NewRepository)
