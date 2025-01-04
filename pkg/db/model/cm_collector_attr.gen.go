package db

import (
	"time"
)

const TableNameCmCollectorAttr = "cm_collector_attr"

// CmCollectorAttr mapped from table <cm_collector_attr>
type CmCollectorAttr struct {
	ID              int64     `gorm:"column:id;primaryKey" json:"id"`
	CollectorTypeID int64     `gorm:"column:collector_type_id;primaryKey" json:"collector_type_id"`
	CollectorID     int64     `gorm:"column:collector_id" json:"collector_id"`
	Key             string    `gorm:"column:key" json:"key"`
	Value           string    `gorm:"column:value" json:"value"`
	Description     string    `gorm:"column:description" json:"description"`
	DeleteYn        string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt           time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt           time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt           time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmCollectorAttr's table name
func (*CmCollectorAttr) TableName() string {
	return TableNameCmCollectorAttr
}
