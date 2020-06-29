package gorm

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"mall/internal/pkg/config"
	"mall/internal/pkg/constant"
)

type Options struct {
	*config.Config
}

// New returns a new gorm.DB instance with options.
func New(opts *Options) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, constant.DatabaseConnectError)
	}

	if err = configure(db, opts); err != nil {
		return nil, errors.Wrap(err, constant.ORMConfigError)
	}

	if opts.EnableAutoMigrate {
		if err = autoMigrate(db); err != nil {
			return nil, errors.Wrap(err, constant.DatabaseMigrateError)
		}
	}

	return db, err
}

// configure gorm.
func configure(db *gorm.DB, opts *Options) error {
	sqlDB, err := db.DB()
	if err != nil {
		return errors.Wrap(err, constant.GetConnectionError)
	}
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	return nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate()
}
