package user

import (
	"gorm.io/plugin/soft_delete"

	"mall/internal/pkg/database/model"
)

// User is the mall user model.
type User struct {
	model.BaseModel
	Username  string                `gorm:"type:varchar(127);not null;uniqueIndex:udx_username;comment:用户名" json:"username"`
	Email     string                `gorm:"type:varchar(127);not null;uniqueIndex:udx_email;comment:邮箱" json:"email"`
	Password  string                `gorm:"type:varchar(255);not null;comment:密码" json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:delete_time;not null;uniqueIndex:udx_username;uniqueIndex:udx_email"`
}

func (u *User) ToSchema() Schema {
	return Schema{
		Username: u.Username,
		Email:    u.Email,
	}
}
