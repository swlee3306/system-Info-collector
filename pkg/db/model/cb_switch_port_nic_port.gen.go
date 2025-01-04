package db

const TableNameCbSwitchPortNicPort = "cb_switch_port_nic_port"

// CbSwitchPortNicPort mapped from table <cb_switch_port_nic_port>
type CbSwitchPortNicPort struct {
	SwitchPortID int64  `gorm:"column:switch_port_id;primaryKey;comment:스위치 포트 아이디" json:"switch_port_id"`                // 스위치 포트 아이디
	NicPortID    int64  `gorm:"column:nic_port_id;primaryKey;comment:닉 포트 아이디" json:"nic_port_id"`                        // 닉 포트 아이디
	DataType     string `gorm:"column:data_type;primaryKey;comment:A = arp(mac-table) L = (LLDP table)" json:"data_type"` // A = arp(mac-table) L = (LLDP table)
}

// TableName CbSwitchPortNicPort's table name
func (*CbSwitchPortNicPort) TableName() string {
	return TableNameCbSwitchPortNicPort
}

