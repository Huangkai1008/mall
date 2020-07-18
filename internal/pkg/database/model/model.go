package model

import (
	"gorm.io/plugin/soft_delete"
	"mall/internal/pkg/util/field"
)

// BaseModel defines custom base model.
// It provides auto createTime/updateTime insert and soft delete.
type BaseModel struct {
	ID        uint                  `gorm:"type:bigint(11) UNSIGNED AUTO_INCREMENT;primaryKey;" json:"id"`
	CreatedAt field.JsonTime        `gorm:"type:datetime;column:create_time;comment:创建时间" json:"-"`
	UpdatedAt field.JsonTime        `gorm:"type:datetime;column:update_time;comment:更新时间" json:"-"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:delete_time;not null;comment:删除时间" json:"-"`
}
