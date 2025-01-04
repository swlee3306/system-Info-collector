package db

import (
	"time"
)

const TableNameCmResourceRelation = "cm_resource_relation"

// CmResourceRelation mapped from table <cm_resource_relation>
type CmResourceRelation struct {
	ID               int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ResourceID       int64     `gorm:"column:resource_id;primaryKey" json:"resource_id"`
	ParentResourceID int64     `gorm:"column:parent_resource_id" json:"parent_resource_id"`
	DeleteYn         string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt            time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt            time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt            time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmResourceRelation's table name
func (*CmResourceRelation) TableName() string {
	return TableNameCmResourceRelation
}
