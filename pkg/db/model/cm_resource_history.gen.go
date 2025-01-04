package db

import (
	"time"
)

const TableNameCmResourceHistory = "cm_resource_history"

// CmResourceHistory mapped from table <cm_resource_history>
type CmResourceHistory struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProviderID     int64     `gorm:"column:provider_id;primaryKey" json:"provider_id"`
	ResourceTypeID int64     `gorm:"column:resource_type_id;primaryKey" json:"resource_type_id"`
	UUID           string    `gorm:"column:uuid" json:"uuid"`
	ResourceName   string    `gorm:"column:resource_name" json:"resource_name"`
	ResourceDesc   string    `gorm:"column:resource_desc" json:"resource_desc"`
	DeleteYn       string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt          time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt          time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt          time.Time `gorm:"column:del_dt" json:"del_dt"`
	EditUser       string    `gorm:"column:edit_user" json:"edit_user"`
	ExecuteKind    string    `gorm:"column:execute_kind" json:"execute_kind"`
	ExecuteDt      time.Time `gorm:"column:execute_dt" json:"execute_dt"`
}

// TableName CmResourceHistory's table name
func (*CmResourceHistory) TableName() string {
	return TableNameCmResourceHistory
}
