package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbCPU = "cb_cpu"

// CbCPU mapped from table <cb_cpu>
type CbCPU struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:cpu 아이디" json:"id"`             // cpu 아이디
	ResourceID  int64          `gorm:"column:resource_id;comment:리소스 아이디" json:"resource_id"`                         // 리소스 아이디
	Type        string         `gorm:"column:type;comment:타입" json:"type"`                                            // 타입
	CPUCore     int64          `gorm:"column:cpu_core;comment:cpu 할당량" json:"cpu_core"`                               // cpu 할당량
	HyperThread bool           `gorm:"column:hyper_thread;comment:하이퍼 쓰레드" json:"hyper_thread"`                       // 하이퍼 쓰레드
	CreatedAt   time.Time      `gorm:"column:created_at;default:current_timestamp();comment:생성 일시" json:"created_at"` // 생성 일시
	UpdatedAt   time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                             // 수정 일시
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                             // 삭제 일시
	Deleted     bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                                   // 삭제 상태
}

// TableName CbCPU's table name
func (*CbCPU) TableName() string {
	return TableNameCbCPU
}
