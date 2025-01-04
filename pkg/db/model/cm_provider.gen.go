package db

import (
	"time"
)

const TableNameCmProvider = "cm_provider"

// CmProvider mapped from table <cm_provider>
type CmProvider struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProviderTypeID int64     `gorm:"column:provider_type_id;primaryKey" json:"provider_type_id"`
	Name           string    `gorm:"column:name" json:"name"`
	Description    string    `gorm:"column:description" json:"description"`
	UUID           string    `gorm:"column:uuid" json:"uuid"`
	DeleteYn       string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt          time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt          time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt          time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmProvider's table name
func (*CmProvider) TableName() string {
	return TableNameCmProvider
}
