package db

const TableNameCbBatonAttr = "cb_baton_attr"

// CbBatonAttr mapped from table <cb_baton_attr>
type CbBatonAttr struct {
	ID    int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Key   string `gorm:"column:key" json:"key"`
	Value string `gorm:"column:value" json:"value"`
}

// TableName CbBatonAttr's table name
func (*CbBatonAttr) TableName() string {
	return TableNameCbBatonAttr
}
