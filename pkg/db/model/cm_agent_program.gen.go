package db

import (
	"time"
)

const TableNameCmAgentProgram = "cm_agent_program"

// CmAgentProgram mapped from table <cm_agent_program>
type CmAgentProgram struct {
	ID                  int64     `gorm:"column:id;primaryKey" json:"id"`
	AgentKindID         int64     `gorm:"column:agent_kind_id;primaryKey" json:"agent_kind_id"`
	AgentProgramVersion string    `gorm:"column:agent_program_version" json:"agent_program_version"`
	AgnetProgramDesc    string    `gorm:"column:agnet_program_desc" json:"agnet_program_desc"`
	IsLatestYn          string    `gorm:"column:is_latest_yn" json:"is_latest_yn"`
	UseYn               string    `gorm:"column:use_yn" json:"use_yn"`
	DelYn               string    `gorm:"column:del_yn" json:"del_yn"`
	RegDt               time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt               time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt               time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmAgentProgram's table name
func (*CmAgentProgram) TableName() string {
	return TableNameCmAgentProgram
}
