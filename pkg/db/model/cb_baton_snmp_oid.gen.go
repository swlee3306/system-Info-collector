package db

const TableNameCbBatonSnmpOid = "cb_baton_snmp_oid"

// CbBatonSnmpOid mapped from table <cb_baton_snmp_oid>
type CbBatonSnmpOid struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ResourceType string `gorm:"column:resource_type" json:"resource_type"`
	ResourceID   int64  `gorm:"column:resource_id;not null" json:"resource_id"`
	OidName      string `gorm:"column:oid_name" json:"oid_name"`
	OidCode      string `gorm:"column:oid_code" json:"oid_code"`
}

// TableName CbBatonSnmpOid's table name
func (*CbBatonSnmpOid) TableName() string {
	return TableNameCbBatonSnmpOid
}
