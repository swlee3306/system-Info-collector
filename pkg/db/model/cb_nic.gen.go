package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbNic = "cb_nic"

// CbNic mapped from table <cb_nic>
type CbNic struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:닉 아이디" json:"id"` // 닉 아이디
	HostID      int64          `gorm:"column:host_id;not null;comment:호스트 아이디" json:"host_id"`          // 호스트 아이디
	Name        string         `gorm:"column:name;comment:닉 이름" json:"name"`                            // 닉 이름
	Type        bool           `gorm:"column:type;comment:타입" json:"type"`                              // 타입
	NetworkType string         `gorm:"column:network_type;comment:네트워크 타입" json:"network_type"`         // 네트워크 타입
	IP          string         `gorm:"column:ip;comment:ip" json:"ip"`                                  // ip
	MacAddr     string         `gorm:"column:mac_addr;comment:맥 주소" json:"mac_addr"`                    // 맥 주소
	CreatedAt   time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"`               // 생성 일시
	UpdatedAt   time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`               // 수정 일시
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`               // 삭제 일시
	Deleted     bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                     // 삭제 상태
}

// TableName CbNic's table name
func (*CbNic) TableName() string {
	return TableNameCbNic
}
