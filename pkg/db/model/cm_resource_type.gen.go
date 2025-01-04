package db

import (
	"time"
)

const TableNameCmResourceType = "cm_resource_type"

// CmResourceType mapped from table <cm_resource_type>
type CmResourceType struct {
	ID                 int64     `gorm:"column:id;primaryKey" json:"id"`
	ResourceTypeKindID int64     `gorm:"column:resource_type_kind_id;primaryKey" json:"resource_type_kind_id"`
	Name               string    `gorm:"column:name" json:"name"`
	Description        string    `gorm:"column:description" json:"description"`
	DeleteYn           string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt              time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt              time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt              time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmResourceType's table name
func (*CmResourceType) TableName() string {
	return TableNameCmResourceType
}
