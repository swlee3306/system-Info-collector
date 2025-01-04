package db

const TableNameCmResource = "cm_resource"

// CmResource mapped from table <cm_resource>
type CmResource struct {
	ID             int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProviderID     int64  `gorm:"column:provider_id;primaryKey" json:"provider_id"`
	ResourceTypeID int64  `gorm:"column:resource_type_id;primaryKey" json:"resource_type_id"`
	UUID           string `gorm:"column:uuid" json:"uuid"`
	ResourceName   string `gorm:"column:resource_name" json:"resource_name"`
	ResourceDesc   string `gorm:"column:resource_desc" json:"resource_desc"`
}

// TableName CmResource's table name
func (*CmResource) TableName() string {
	return TableNameCmResource
}
