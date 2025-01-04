package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbMemory = "cb_memory"

// CbMemory mapped from table <cb_memory>
type CbMemory struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:memory 아이디" json:"id"`          // memory 아이디
	ResourceID  int64          `gorm:"column:resource_id;comment:리소스 아이디" json:"resource_id"`                         // 리소스 아이디
	Type        string         `gorm:"column:type;comment:타입" json:"type"`                                            // 타입
	MemoryTotal int64          `gorm:"column:memory_total;comment:메모리 할당량" json:"memory_total"`                       // 메모리 할당량
	CreatedAt   time.Time      `gorm:"column:created_at;default:current_timestamp();comment:생성 일시" json:"created_at"` // 생성 일시
	UpdatedAt   time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                             // 수정 일시
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                             // 삭제 일시
	Deleted     bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                                   // 삭제 상태
}

// TableName CbMemory's table name
func (*CbMemory) TableName() string {
	return TableNameCbMemory
}
