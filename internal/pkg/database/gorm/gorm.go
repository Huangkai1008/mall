package gorm

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"mall/internal/app/v1/product"
	"mall/internal/app/v1/user"
	"mall/internal/pkg/constant"
)

// New returns a new gorm.DB instance with options.
func New(o *Options) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.DSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, constant.DatabaseConnectError)
	}

	if err = configure(db, o); err != nil {
		return nil, errors.Wrap(err, constant.ORMConfigError)
	}

	if o.EnableAutoMigrate {
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
	return db.AutoMigrate(
		&user.User{},
		&product.Brand{}, &product.Store{}, &product.Category{},
		&product.Spu{},
	)
}

var ProviderSet = wire.NewSet(New, NewOptions)
