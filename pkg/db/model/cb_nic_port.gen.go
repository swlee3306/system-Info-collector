package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbNicPort = "cb_nic_port"

// CbNicPort mapped from table <cb_nic_port>
type CbNicPort struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:닉 포트 아이디" json:"id"` // 닉 포트 아이디
	NicID     int64          `gorm:"column:nic_id;not null;comment:닉 아이디" json:"nic_id"`                 // 닉 아이디
	CreatedAt time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"`                  // 생성 일시
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                  // 수정 일시
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                  // 삭제 일시
	Deleted   bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                        // 삭제 상태
	Type      string         `gorm:"column:type;comment:타입" json:"type"`                                 // 타입
	IP        string         `gorm:"column:ip;comment:ip" json:"ip"`                                     // ip
	MacAddr   string         `gorm:"column:mac_addr;comment:mac 주소" json:"mac_addr"`                     // mac 주소
}

// TableName CbNicPort's table name
func (*CbNicPort) TableName() string {
	return TableNameCbNicPort
}
