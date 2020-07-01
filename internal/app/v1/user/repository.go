package user

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	repo "mall/internal/pkg/repository"
)

type Repository interface {
	Exist(condition interface{}) (bool, error)
}

type repository struct {
	*repo.GormRepository
}

// NewRepository returns new user Repository.
func NewRepository(logger *zap.Logger, db *gorm.DB) Repository {
	return &repository{
		GormRepository: &repo.GormRepository{
			Logger: logger.With(zap.String("type", "UserRepository")),
			Db:     db,
		}}
}
