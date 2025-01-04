package db

const TableNameCbVMNetworkVlan = "cb_vm_network_vlan"

// CbVMNetworkVlan mapped from table <cb_vm_network_vlan>
type CbVMNetworkVlan struct {
	ID            int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProviderID             int64          `gorm:"column:provider_id;not null"`
	VMUUID        string `gorm:"column:vm_uuid;not null" json:"vm_uuid"`
	NetworkVlanUUID string `gorm:"column:network_vlan_uuid;not null" json:"network_vlan_uuid"`
}

// TableName CbVMNetworkVlan's table name
func (*CbVMNetworkVlan) TableName() string {
	return TableNameCbVMNetworkVlan
}
