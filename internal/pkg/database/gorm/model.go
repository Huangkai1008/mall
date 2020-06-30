package gorm

import (
	"gorm.io/plugin/soft_delete"

	"mall/internal/pkg/util"
)

// BaseModel defines custom base model.
// It provides auto createTime/updateTime insert and soft delete.
type BaseModel struct {
	ID        uint                  `gorm:"type:bigint(11) unsigned;autoIncrement;primaryKey" json:"id"`
	CreatedAt util.JsonTime         `gorm:"type:datetime;column:create_time;comment:'创建时间'" json:"create_time"`
	UpdatedAt util.JsonTime         `gorm:"type:datetime;column:update_time;comment:'更新时间'" json:"update_time"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:delete_time"`
}
