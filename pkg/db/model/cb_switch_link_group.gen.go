package db

const TableNameCbSwitchLinkGroup = "cb_switch_link_group"

// CbSwitchLinkGroup mapped from table <cb_switch_link_group>
type CbSwitchLinkGroup struct {
	ID               int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"` // id
	SwitchID         int64  `gorm:"column:switch_id;comment:스위치 id" json:"switch_id"`             // 스위치 id
	Type             string `gorm:"column:type;comment:타입" json:"type"`                           // 타입
	Master           string `gorm:"column:master;comment:master/slave" json:"master"`             // master/slave
	SwitchPortFilter string `gorm:"column:switch_port_filter" json:"switch_port_filter"`
	Group_           int64  `gorm:"column:group;not null;comment:스태킹 및 본딩으로 묶여있는 그룹" json:"group"` // 스태킹 및 본딩으로 묶여있는 그룹
}

// TableName CbSwitchLinkGroup's table name
func (*CbSwitchLinkGroup) TableName() string {
	return TableNameCbSwitchLinkGroup
}
