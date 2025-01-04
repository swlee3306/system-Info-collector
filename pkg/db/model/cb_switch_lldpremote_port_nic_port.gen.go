package db

const TableNameCbSwitchLldpRemotePortNicPort = "cb_switch_lldpRemote_port_nic_port"

// CbSwitchLldpRemotePortNicPort mapped from table <cb_switch_lldpRemote_port_nic_port>
type CbSwitchLldpRemotePortNicPort struct {
	SwitchesID   int64 `gorm:"column:switches_id;primaryKey;comment:스위치 리소스 아이디" json:"switches_id"`     // 스위치 리소스 아이디
	SwitchPortNo int64 `gorm:"column:switch_port_no;primaryKey;comment:스위치 포트 넘버" json:"switch_port_no"` // 스위치 포트 넘버
	NicPortID    int64 `gorm:"column:nic_port_id;primaryKey;comment:닉 포트 아이디" json:"nic_port_id"`        // 닉 포트 아이디
}

// TableName CbSwitchLldpRemotePortNicPort's table name
func (*CbSwitchLldpRemotePortNicPort) TableName() string {
	return TableNameCbSwitchLldpRemotePortNicPort
}

