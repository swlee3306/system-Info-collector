package db

const TableNameCbCephHost = "cb_ceph_host"

// CbCephHost mapped from table <cb_ceph_host>
type CbCephHost struct {
	ID     int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RackID int64  `gorm:"column:rack_id;comment:렉 아이디" json:"rack_id"` // 렉 아이디
	Key    string `gorm:"column:key" json:"key"`
	Value  string `gorm:"column:value" json:"value"`
}

// TableName CbCephHost's table name
func (*CbCephHost) TableName() string {
	return TableNameCbCephHost
}
