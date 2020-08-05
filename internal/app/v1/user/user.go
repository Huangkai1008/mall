package user

import (
	"gorm.io/plugin/soft_delete"
	metav1 "mall/pkg/meta/v1"
)

// User represents a user restful resource.
type User struct {
	metav1.ObjectMeta

	// Required: true
	Username string `json:"username" gorm:"type:varchar(127);not null;uniqueIndex:udx_username;comment:用户名"`

	// Required: true
	Email string `json:"email" gorm:"type:varchar(127);not null;uniqueIndex:udx_email;comment:邮箱"`

	// Required: true
	Password string `json:"-" gorm:"type:varchar(255);not null;comment:密码"`

	// Required: true
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"column:delete_time;not null;uniqueIndex:udx_username;uniqueIndex:udx_email;comment:删除时间"`
}
