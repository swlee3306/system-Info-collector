package db

const TableNameCmResourceTmp = "cm_resource_tmp"

// CmResourceTmp mapped from table <cm_resource_tmp>
type CmResourceTmp struct {
	ID             int64  `gorm:"column:id;not null" json:"id"`
	ProviderID     int64  `gorm:"column:provider_id;not null" json:"provider_id"`
	ResourceTypeID int64  `gorm:"column:resource_type_id;not null" json:"resource_type_id"`
	UUID           string `gorm:"column:uuid" json:"uuid"`
	ResourceName   string `gorm:"column:resource_name" json:"resource_name"`
	ResourceDesc   string `gorm:"column:resource_desc" json:"resource_desc"`
}

// TableName CmResourceTmp's table name
func (*CmResourceTmp) TableName() string {
	return TableNameCmResourceTmp
}
