package db

import (
	"time"
)

const TableNameCmCollectorType = "cm_collector_type"

// CmCollectorType mapped from table <cm_collector_type>
type CmCollectorType struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	DeleteYn    string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt       time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt       time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt       time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmCollectorType's table name
func (*CmCollectorType) TableName() string {
	return TableNameCmCollectorType
}
