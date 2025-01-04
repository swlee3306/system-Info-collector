package db

const TableNameCbResourceCfg = "cb_resource_cfg"

// CbResourceCfg mapped from table <cb_resource_cfg>
type CbResourceCfg struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ResourceType string `gorm:"column:resource_type" json:"resource_type"`
	ResourceID   int64  `gorm:"column:resource_id" json:"resource_id"`
	Key          string `gorm:"column:key" json:"key"`
	Value        string `gorm:"column:value" json:"value"`
}

// TableName CbResourceCfg's table name
func (*CbResourceCfg) TableName() string {
	return TableNameCbResourceCfg
}
