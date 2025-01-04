package db

import (
	"time"
)

const TableNameCmAgentKind = "cm_agent_kind"

// CmAgentKind mapped from table <cm_agent_kind>
type CmAgentKind struct {
	ID       int64     `gorm:"column:id;primaryKey" json:"id"`
	KindName string    `gorm:"column:kind_name" json:"kind_name"`
	KindDesc string    `gorm:"column:kind_desc" json:"kind_desc"`
	RegDt    time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt    time.Time `gorm:"column:mod_dt" json:"mod_dt"`
}

// TableName CmAgentKind's table name
func (*CmAgentKind) TableName() string {
	return TableNameCmAgentKind
}
