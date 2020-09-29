package database

import (
	"gorm.io/gorm"

	orm "github.com/Huangkai1008/micro-kit/pkg/database/gorm"

	"github.com/Huangkai1008/mall/internal/pkg/config"
)

func NewGorm(c *config.Config, tables []interface{}) (*gorm.DB, error) {
	return orm.New(
		tables,
		orm.WithSource(c.Database.Source),
	)
}
