package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GormRepository struct {
	Logger *zap.Logger // repo logger
	Db     *gorm.DB    // repo db connection
}

func (r *GormRepository) GetAll(conditions interface{}) (records []*interface{}, err error) {
	err = r.Db.Where(conditions).Find(&records).Error
	return
}

func (r *GormRepository) GetOne(conditions interface{}) (record *interface{}, err error) {
	err = r.Db.Where(conditions).First(&record).Error
	return
}

func (r *GormRepository) Exist(conditions interface{}) (bool, error) {
	var (
		err    error
		record interface{}
	)

	err = r.Db.Where(conditions).Limit(1).Find(&record).Error
	if err != nil {
		return false, err
	} else {
		return record == nil, err
	}
}

func (r *GormRepository) Create(record interface{}) error {
	err := r.Db.Create(&record).Error
	return err
}
