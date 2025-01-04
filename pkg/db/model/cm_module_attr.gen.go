package db

import (
	"time"
)

const TableNameCmModuleAttr = "cm_module_attr"

// CmModuleAttr mapped from table <cm_module_attr>
type CmModuleAttr struct {
	ID          int64     `gorm:"column:id;not null" json:"id"`
	ModuleID    int64     `gorm:"column:module_id;not null" json:"module_id"`
	Key         string    `gorm:"column:key" json:"key"`
	Value       string    `gorm:"column:value" json:"value"`
	Description string    `gorm:"column:description" json:"description"`
	DeleteYn    string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt       time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt       time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt       time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmModuleAttr's table name
func (*CmModuleAttr) TableName() string {
	return TableNameCmModuleAttr
}
