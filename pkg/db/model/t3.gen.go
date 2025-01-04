package db

const TableNameT3 = "t3"

// T3 mapped from table <t3>
type T3 struct {
	C1 int32 `gorm:"column:c1;primaryKey;autoIncrement:true" json:"c1"`
	C2 int32 `gorm:"column:c2" json:"c2"`
}

// TableName T3's table name
func (*T3) TableName() string {
	return TableNameT3
}
