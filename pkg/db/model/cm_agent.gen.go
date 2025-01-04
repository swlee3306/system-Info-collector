package db

import (
	"time"
)

const TableNameCmAgent = "cm_agent"

// CmAgent mapped from table <cm_agent>
type CmAgent struct {
	ID               int64     `gorm:"column:id;primaryKey" json:"id"`
	RespirceID       int64     `gorm:"column:respirce_id;primaryKey" json:"respirce_id"`
	AgentProgramID   int64     `gorm:"column:agent_program_id;primaryKey" json:"agent_program_id"`
	AgentKindID      int64     `gorm:"column:agent_kind_id;primaryKey" json:"agent_kind_id"`
	AgnetStatus      string    `gorm:"column:agnet_status" json:"agnet_status"`
	AgentStatusPrevv string    `gorm:"column:agent_status_prevv" json:"agent_status_prevv"`
	DeleteYn         string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt            time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt            time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt            time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmAgent's table name
func (*CmAgent) TableName() string {
	return TableNameCmAgent
}
