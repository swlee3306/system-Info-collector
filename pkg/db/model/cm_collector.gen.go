package db

import (
	"time"
)

const TableNameCmCollector = "cm_collector"

// CmCollector mapped from table <cm_collector>
type CmCollector struct {
	ID              int64     `gorm:"column:id;primaryKey" json:"id"`
	CollectorTypeID int64     `gorm:"column:collector_type_id;primaryKey" json:"collector_type_id"`
	Name            string    `gorm:"column:name" json:"name"`
	Description     string    `gorm:"column:description" json:"description"`
	DeleteYn        string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegGt           time.Time `gorm:"column:reg_gt" json:"reg_gt"`
	ModDt           time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt           time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmCollector's table name
func (*CmCollector) TableName() string {
	return TableNameCmCollector
}
