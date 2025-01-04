package db

const TableNameCbResourceAttr = "cb_resource_attr"

// CbResourceAttr mapped from table <cb_resource_attr>
type CbResourceAttr struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ResourceType string `gorm:"column:resource_type" json:"resource_type"`
	ResourceID   int64  `gorm:"column:resource_id" json:"resource_id"`
	Key          string `gorm:"column:key" json:"key"`
	Value        string `gorm:"column:value" json:"value"`
}

// TableName CbResourceAttr's table name
func (*CbResourceAttr) TableName() string {
	return TableNameCbResourceAttr
}
