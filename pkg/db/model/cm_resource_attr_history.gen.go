package db

import (
	"time"
)

const TableNameCmResourceAttrHistory = "cm_resource_attr_history"

// CmResourceAttrHistory mapped from table <cm_resource_attr_history>
type CmResourceAttrHistory struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID        string    `gorm:"column:uuid;primaryKey" json:"uuid"`
	Key         string    `gorm:"column:key" json:"key"`
	Value       string    `gorm:"column:value" json:"value"`
	Description string    `gorm:"column:description" json:"description"`
	ValueType   string    `gorm:"column:value_type" json:"value_type"`
	DeleteYn    string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt       time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt       time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt       time.Time `gorm:"column:del_dt" json:"del_dt"`
	EditUser    string    `gorm:"column:edit_user" json:"edit_user"`
	ExecuteKind string    `gorm:"column:execute_kind" json:"execute_kind"`
	ExecuteDt   time.Time `gorm:"column:execute_dt" json:"execute_dt"`
}

// TableName CmResourceAttrHistory's table name
func (*CmResourceAttrHistory) TableName() string {
	return TableNameCmResourceAttrHistory
}
