package db

const TableNameT1 = "t1"

// T1 mapped from table <t1>
type T1 struct {
	C1 string `gorm:"column:c1" json:"c1"`
}

// TableName T1's table name
func (*T1) TableName() string {
	return TableNameT1
}
