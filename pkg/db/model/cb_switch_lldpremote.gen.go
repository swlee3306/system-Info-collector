package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbSwitchLldpRemote = "cb_switch_lldpRemote"

// CbSwitchLldpRemote mapped from table <cb_switch_lldpRemote>
type CbSwitchLldpRemote struct {
	ID                      int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:lldp table 아이디" json:"id"`                            // lldp table 아이디
	SwitchesID              int64          `gorm:"column:switches_id;not null;comment:스위치 아이디" json:"switches_id"`                                      // 스위치 아이디
	SwitchPortNo            int64          `gorm:"column:switch_port_no;not null;comment:스위치 포트 넘버" json:"switch_port_no"`                              // 스위치 포트 넘버
	LldpRemChassisIDSubtype int64          `gorm:"column:lldpRemChassisIdSubtype;comment:lldp chassisIdSubtype" json:"lldpRemChassisIdSubtype"`         // lldp chassisIdSubtype
	LldpRemChassisID        string         `gorm:"column:lldpRemChassisId;comment:lldp ChassisId" json:"lldpRemChassisId"`                              // lldp ChassisId
	LldpRemPortIDSubtype    int64          `gorm:"column:lldpRemPortIdSubtype;comment:lldp PortIdSubtype" json:"lldpRemPortIdSubtype"`                  // lldp PortIdSubtype
	LldpRemPortID           string         `gorm:"column:lldpRemPortId;comment:lldp PortId" json:"lldpRemPortId"`                                       // lldp PortId
	LldpRemPortDesc         string         `gorm:"column:lldpRemPortDesc;comment:lldp RemotePortDesc" json:"lldpRemPortDesc"`                           // lldp RemotePortDesc
	LldpRemSysName          string         `gorm:"column:lldpRemSysName;comment:lldp RemotePort system-name" json:"lldpRemSysName"`                     // lldp RemotePort system-name
	LldpRemSysDesc          string         `gorm:"column:lldpRemSysDesc;comment:lldp Remote System-Description" json:"lldpRemSysDesc"`                  // lldp Remote System-Description
	LldpRemSysCapSupported  string         `gorm:"column:lldpRemSysCapSupported;comment:lldp Remote System-CapSupported" json:"lldpRemSysCapSupported"` // lldp Remote System-CapSupported
	LldpRemSysCapEnabled    string         `gorm:"column:lldpRemSysCapEnabled;comment:lldp Remote System-CapEnabled" json:"lldpRemSysCapEnabled"`       // lldp Remote System-CapEnabled
	CreatedAt               time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"`                                                   // 생성 일시
	UpdatedAt               time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                                                   // 수정 일시
	DeletedAt               gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                                                   // 삭제 일시
}

// TableName CbSwitchLldpRemote's table name
func (*CbSwitchLldpRemote) TableName() string {
	return TableNameCbSwitchLldpRemote
}
