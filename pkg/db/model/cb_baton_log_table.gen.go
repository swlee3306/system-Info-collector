package db

import (
	"time"
)

const TableNameCbBatonLogTable = "cb_baton_log_table"

// CbBatonLogTable mapped from table <cb_baton_log_table>
type CbBatonLogTable struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	LogLevel   string    `gorm:"column:log_level;not null" json:"log_level"`
	Target     string    `gorm:"column:target;not null" json:"target"`
	Detail     string    `gorm:"column:detail" json:"detail"`
	LogMessage string    `gorm:"column:log_message;not null" json:"log_message"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	InsertAt   time.Time `gorm:"column:insert_at;default:current_timestamp()" json:"insert_at"`
}

// TableName CbBatonLogTable's table name
func (*CbBatonLogTable) TableName() string {
	return TableNameCbBatonLogTable
}
