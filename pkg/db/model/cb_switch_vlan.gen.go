package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbSwitchVlan = "cb_switch_vlan"

// CbSwitchVlan mapped from table <cb_switch_vlan>
type CbSwitchVlan struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SwitchesID int64          `gorm:"column:switches_id;not null;comment:스위치 아이디" json:"switches_id"`                                          // 스위치 아이디
	Status     string         `gorm:"column:status;comment:active, notInService, notReady, createAndGo, createAndWait, destroy" json:"status"` // active, notInService, notReady, createAndGo, createAndWait, destroy
	VlanID     int64          `gorm:"column:vlan_id;comment:스위치 vlan 아이디" json:"vlan_id"`                                                      // 스위치 vlan 아이디
	VlanName   string         `gorm:"column:vlan_name;comment:스위치 vlan 이름" json:"vlan_name"`                                                   // 스위치 vlan 이름
	CreatedAt  time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"`                                                       // 생성 일시
	UpdatedAt  time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                                                       // 수정 일시
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                                                       // 삭제 일시
	Deleted    bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                                                             // 삭제 상태
}

// TableName CbSwitchVlan's table name
func (*CbSwitchVlan) TableName() string {
	return TableNameCbSwitchVlan
}
