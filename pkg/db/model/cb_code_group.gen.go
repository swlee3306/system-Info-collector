package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbCodeGroup = "cb_code_group"

// CbCodeGroup mapped from table <cb_code_group>
type CbCodeGroup struct {
	Code        string         `gorm:"column:code;primaryKey" json:"code"`
	Name        string         `gorm:"column:name" json:"name"`
	Description string         `gorm:"column:description" json:"description"`
	CreatedAt   time.Time      `gorm:"column:created_at;default:current_timestamp();comment:생성 일시" json:"created_at"` // 생성 일시
	UpdatedAt   time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                             // 수정 일시
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                             // 삭제 일시
	Deleted     bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                                   // 삭제 상태
}

// TableName CbCodeGroup's table name
func (*CbCodeGroup) TableName() string {
	return TableNameCbCodeGroup
}
