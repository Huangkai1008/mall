package account

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"

	e "github.com/Huangkai1008/micro-kit/pkg/error"
	metav1 "github.com/Huangkai1008/micro-kit/pkg/meta/v1"

	"github.com/Huangkai1008/mall/internal/pkg/constant"
)

// Account represents a account restful resource.
type Account struct {
	metav1.ObjectMeta

	// Required: true
	Username string `json:"username" database:"type:varchar(127);not null;uniqueIndex:udx_username;comment:用户名"`

	// Required: true
	Email string `json:"email" database:"type:varchar(127);not null;uniqueIndex:udx_email;comment:邮箱"`

	// Required: true
	Password string `json:"-" database:"type:varchar(255);not null;comment:密码"`

	// Required: true
	DeletedAt soft_delete.DeletedAt `json:"-" database:"column:delete_time;not null;uniqueIndex:udx_username;uniqueIndex:udx_email;comment:删除时间"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if r := tx.Where("username = ?", a.Username).Limit(1).Find(&Account{}); r.Error != nil {
		return r.Error
	} else if r.RowsAffected > 0 {
		return e.NewBadRequestError(constant.AccountAlreadyExist)
	}

	if r := tx.Where("email = ?", a.Email).Limit(1).Find(&Account{}); r.Error != nil {
		return r.Error
	} else if r.RowsAffected > 0 {
		return e.NewBadRequestError(constant.AccountEmailAlreadyExist)
	}
	return nil
}
