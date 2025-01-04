package db

const TableNameCmResourceAttr = "cm_resource_attr"

// CmResourceAttr mapped from table <cm_resource_attr>
type CmResourceAttr struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ResourceID  int64  `gorm:"column:resource_id;primaryKey" json:"resource_id"`
	Key         string `gorm:"column:key" json:"key"`
	Value       string `gorm:"column:value" json:"value"`
	Description string `gorm:"column:description" json:"description"`
	ValueType   string `gorm:"column:value_type" json:"value_type"`
}

// TableName CmResourceAttr's table name
func (*CmResourceAttr) TableName() string {
	return TableNameCmResourceAttr
}
