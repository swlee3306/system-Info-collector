package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbDisk = "cb_disk"

// CbDisk mapped from table <cb_disk>
type CbDisk struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:disk 아이디" json:"id"`            // disk 아이디
	ResourceID int64          `gorm:"column:resource_id;comment:리소스 아이디" json:"resource_id"`                         // 리소스 아이디
	Type       string         `gorm:"column:type;comment:타입" json:"type"`                                            // 타입
	DiskTotal  int64          `gorm:"column:disk_total;comment:디스크 할당량" json:"disk_total"`                           // 디스크 할당량
	CreatedAt  time.Time      `gorm:"column:created_at;default:current_timestamp();comment:생성 일시" json:"created_at"` // 생성 일시
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                             // 수정 일시
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                             // 삭제 일시
	Deleted    bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                                   // 삭제 상태
	DiskType   string         `gorm:"column:disk_type;comment:디스크 타입" json:"disk_type"`                              // 디스크 타입
	Lvm        string         `gorm:"column:lvm;comment:lvm" json:"lvm"`                                             // lvm
}

// TableName CbDisk's table name
func (*CbDisk) TableName() string {
	return TableNameCbDisk
}
