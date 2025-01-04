package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbRackCenter = "cb_rack_center"

// CbRackCenter mapped from table <cb_rack_center>
type CbRackCenter struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:렉 센터 아이디" json:"id"` // 렉 센터 아이디
	UUID        string         `gorm:"column:uuid;not null;default:uuid();comment:렉 센터 UUID" json:"uuid"`  // 렉 센터 UUID
	Name        string         `gorm:"column:name;comment:렉 센터 이름" json:"name"`                            // 렉 센터 이름
	Location    string         `gorm:"column:location;comment:렉 위치" json:"location"`                       // 렉 위치
	RackCnt     int32          `gorm:"column:rack_cnt;comment:렉 센터 내 렉의 수량" json:"rack_cnt"`               // 렉 센터 내 렉의 수량
	Status      string         `gorm:"column:status;comment:렉 상태" json:"status"`                           // 렉 상태
	Description string         `gorm:"column:description;comment:설명" json:"description"`                   // 설명
	Usable      string         `gorm:"column:usable;comment:용도" json:"usable"`                             // 용도
	CreatedAt   time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"`                  // 생성 일시
	UpdatedAt   time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                  // 수정 일시
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                  // 삭제 일시
	Deleted     bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                        // 삭제 상태
}

// TableName CbRackCenter's table name
func (*CbRackCenter) TableName() string {
	return TableNameCbRackCenter
}
