package db

import (
	"time"
)

const TableNameCmModule = "cm_module"

// CmModule mapped from table <cm_module>
type CmModule struct {
	ID           int64     `gorm:"column:id;primaryKey" json:"id"`
	ProviderID   int64     `gorm:"column:provider_id;primaryKey" json:"provider_id"`
	ModuleTypeID int64     `gorm:"column:module_type_id;primaryKey" json:"module_type_id"`
	Name         string    `gorm:"column:name" json:"name"`
	Description  string    `gorm:"column:description" json:"description"`
	DeleteYn     string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt        time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt        time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt        time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmModule's table name
func (*CmModule) TableName() string {
	return TableNameCmModule
}
