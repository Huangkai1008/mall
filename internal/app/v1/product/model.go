package product

import (
	"gorm.io/plugin/soft_delete"

	"mall/internal/pkg/database/model"
)

// Brand is the product brand model.
type Brand struct {
	model.BaseModel
	BrandName string                `gorm:"varchar(64);not null;uniqueIndex:udx_brand_name;comment:品牌名称" json:"brand_name"`
	BrandLogo string                `gorm:"varchar(255);comment:品牌logo" json:"brand_logo"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:delete_time;not null;uniqueIndex:udx_brand_name;comment:删除时间" json:"-"`
}

func (Brand) TableName() string {
	return "product_brand"
}

// Store is the product store model.
type Store struct {
	model.BaseModel
	StoreName string `gorm:"varchar(64);not null;uniqueIndex:udx_store_name;comment:商铺名称" json:"store_name"`
	StoreLogo string `gorm:"varchar(255);comment:商铺logo" json:"store_logo"`
	StoreDesc string `gorm:"varchar(255);comment:商铺描述" json:"store_desc"`
}

func (Store) TableName() string {
	return "product_store"
}

// Category is the product category model.
type Category struct {
	model.BaseModel
	ParentId    uint                  `gorm:"type:bigint(11) UNSIGNED;not null;index;comment:父分类ID，0表示1级分类" json:"parent_id"`
	CatName     string                `gorm:"varchar(64);not null;uniqueIndex:udx_cat_name;comment:分类名称" json:"cat_name"`
	CatLevel    uint8                 `gorm:"type:tinyint(1);not null;index;comment:分类等级：0->1级,1->2级" json:"cat_level"`
	CatKeywords string                `gorm:"type:varchar(255);comment:分类关键词" json:"cat_keywords"`
	CatIcon     string                `gorm:"type:varchar(255);comment:分类图标" json:"cat_icon"`
	CatDesc     string                `gorm:"type:varchar(255);comment:分类描述" json:"cat_desc"`
	DeletedAt   soft_delete.DeletedAt `gorm:"column:delete_time;not null;uniqueIndex:udx_cat_name;comment:删除时间" json:"-"`
}

func (Category) TableName() string {
	return "product_category"
}
