package product

//
//import (
//	"database.io/plugin/soft_delete"
//
//	"mall/internal/pkg/database/model"
//)
//
//// Brand is the product brand model.
//type Brand struct {
//	model.BaseModel
//	BrandName string                `database:"type:varchar(64);not null;uniqueIndex:udx_brand_name;comment:品牌名称" json:"brand_name"`
//	BrandLogo string                `database:"type:varchar(255);comment:品牌logo" json:"brand_logo"`
//	DeletedAt soft_delete.DeletedAt `database:"column:delete_time;not null;uniqueIndex:udx_brand_name;comment:删除时间" json:"-"`
//}
//
//func (Brand) TableName() string {
//	return "product_brand"
//}
//
//// Store is the product store model.
//type Store struct {
//	model.BaseModel
//	StoreName string `database:"type:varchar(64);not null;uniqueIndex:udx_store_name;comment:商铺名称" json:"store_name"`
//	StoreLogo string `database:"type:varchar(255);comment:商铺logo" json:"store_logo"`
//	StoreDesc string `database:"type:varchar(255);comment:商铺描述" json:"store_desc"`
//}
//
//func (Store) TableName() string {
//	return "product_store"
//}
//
//// Category is the product category model.
//type Category struct {
//	model.BaseModel
//	ParentId    uint                  `database:"type:bigint(11) UNSIGNED;not null;index;comment:父分类ID，0表示1级分类" json:"parent_id"`
//	CatName     string                `database:"type:varchar(64);not null;uniqueIndex:udx_cat_name;comment:分类名称" json:"cat_name"`
//	CatLevel    uint8                 `database:"type:tinyint(1);not null;index;comment:分类等级：0->1级,1->2级" json:"cat_level"`
//	CatKeywords string                `database:"type:varchar(255);comment:分类关键词" json:"cat_keywords"`
//	CatIcon     string                `database:"type:varchar(255);comment:分类图标" json:"cat_icon"`
//	CatDesc     string                `database:"type:varchar(255);comment:分类描述" json:"cat_desc"`
//	DeletedAt   soft_delete.DeletedAt `database:"column:delete_time;not null;uniqueIndex:udx_cat_name;comment:删除时间" json:"-"`
//}
//
//func (Category) TableName() string {
//	return "product_category"
//}
//
//// Spu is the product spu model.
//type Spu struct {
//	model.BaseModel
//	SpuName   string                `database:"type:varchar(64);not null;index;comment:商品名称" json:"spu_name"`
//	SpuSn     string                `database:"type:varchar(24);not null;uniqueIndex:udx_spu_sn;comment:商品货号" json:"spu_sn"`
//	SubTitle  string                `database:"type:varchar(127);comment:副标题" json:"sub_title"`
//	BrandId   uint                  `database:"type:bigint(11) UNSIGNED;not null;index;comment:商品品牌id" json:"brand_id"`
//	BrandName string                `database:"type:varchar(64);not null;comment:品牌名称" json:"brand_name"`
//	StoreId   uint                  `database:"type:bigint(11) UNSIGNED;not null;index;comment:商品店铺id" json:"store_id"`
//	StoreName string                `database:"type:varchar(64);not null;comment:商铺名称" json:"store_name"`
//	CatId     uint                  `database:"type:bigint(11) UNSIGNED;not null;index;comment:商品分类id" json:"cat_id"`
//	CatName   string                `database:"type:varchar(64);not null;comment:分类名称" json:"cat_name"`
//	Unit      string                `database:"type:varchar(32);comment:单位：件/台..." json:"unit"`
//	Published *bool                 `database:"type:tinyint(1);not null;index;comment:上架状态" json:"published"`
//	DeletedAt soft_delete.DeletedAt `database:"column:delete_time;not null;uniqueIndex:udx_spu_sn;comment:删除时间" json:"-"`
//}
//
//func (Spu) TableName() string {
//	return "product_spu"
//}
