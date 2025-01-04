package db

const TableNameCbSwitchPortMacIP = "cb_switch_port_mac_ip"

// CbSwitchPortMacIP mapped from table <cb_switch_port_mac_ip>
type CbSwitchPortMacIP struct {
	SwitchPortID int64  `gorm:"column:switch_port_id;primaryKey" json:"switch_port_id"`
	IPAddress    string `gorm:"column:ip_address;primaryKey" json:"ip_address"`
	MacAddress   string `gorm:"column:mac_address;primaryKey" json:"mac_address"`
}

// TableName CbSwitchPortMacIP's table name
func (*CbSwitchPortMacIP) TableName() string {
	return TableNameCbSwitchPortMacIP
}
