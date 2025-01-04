package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbRackResourcePo = "cb_rack_resource_pos"

// CbRackResourcePo mapped from table <cb_rack_resource_pos>
type CbRackResourcePo struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:렉 리소스 포스 아이디" json:"id"` // 렉 리소스 포스 아이디
	RackID       int64          `gorm:"column:rack_id;not null;comment:렉 아이디" json:"rack_id"`                   // 렉 아이디
	ResourceID   int64          `gorm:"column:resource_id;not null;comment:자원 아이디" json:"resource_id"`          // 자원 아이디
	ResourceType string         `gorm:"column:resource_type;not null;comment:자원 타입" json:"resource_type"`       // 자원 타입
	SlotPos      int32          `gorm:"column:slot_pos;comment:슬롯 시작 위치" json:"slot_pos"`                       // 슬롯 시작 위치
	SlotSize     int32          `gorm:"column:slot_size;comment:슬롯 사이즈" json:"slot_size"`                       // 슬롯 사이즈
	CreatedAt    time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"`                      // 생성 일시
	UpdatedAt    time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                      // 수정 일시
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                      // 삭제 일시
	Deleted      bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                            // 삭제 상태
	Status       string         `gorm:"column:status;comment:상태" json:"status"`                                 // 상태
}

// TableName CbRackResourcePo's table name
func (*CbRackResourcePo) TableName() string {
	return TableNameCbRackResourcePo
}
