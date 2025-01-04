package db

const TableNameCbSwitchPortVlan = "cb_switch_port_vlan"

// CbSwitchPortVlan mapped from table <cb_switch_port_vlan>
type CbSwitchPortVlan struct {
	SwitchPortID                 int64 `gorm:"column:switch_port_id;primaryKey" json:"switch_port_id"`
	SwitchVlanID                 int64 `gorm:"column:switch_vlan_id;primaryKey" json:"switch_vlan_id"`
	SwitchPortSwitchesIDPhysical int64 `gorm:"column:switch_port_switches_id_physical;not null" json:"switch_port_switches_id_physical"`
	SwitchPortPortNo             int64 `gorm:"column:switch_port_port_no;not null" json:"switch_port_port_no"`
	SwitchVlanVlanID             int64 `gorm:"column:switch_vlan_vlan_id;not null" json:"switch_vlan_vlan_id"`
}

// TableName CbSwitchPortVlan's table name
func (*CbSwitchPortVlan) TableName() string {
	return TableNameCbSwitchPortVlan
}
