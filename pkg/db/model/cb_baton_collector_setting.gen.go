package db

const TableNameCbBatonCollectorSetting = "cb_baton_collector_setting"

// CbBatonCollectorSetting mapped from table <cb_baton_collector_setting>
type CbBatonCollectorSetting struct {
	ID                          int32  `gorm:"column:id;primaryKey" json:"id"`
	Name                        string `gorm:"column:name;not null" json:"name"`
	Description                 string `gorm:"column:description" json:"description"`
	CollectorDefPeriodSecMeta   string `gorm:"column:collector_defPeriodSecMeta;default:60" json:"collector_defPeriodSecMeta"`
	CollectorDefPeriodSecMetric string `gorm:"column:collector_defPeriodSecMetric;default:60" json:"collector_defPeriodSecMetric"`
	OutEndpoints                string `gorm:"column:OutEndpoints" json:"OutEndpoints"`
}

// TableName CbBatonCollectorSetting's table name
func (*CbBatonCollectorSetting) TableName() string {
	return TableNameCbBatonCollectorSetting
}
