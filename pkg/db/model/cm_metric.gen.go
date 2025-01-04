package db

import (
	"time"
)

const TableNameCmMetric = "cm_metric"

// CmMetric mapped from table <cm_metric>
type CmMetric struct {
	ID               int64     `gorm:"column:id;primaryKey" json:"id"`
	ProviderTypeID   int64     `gorm:"column:provider_type_id;primaryKey" json:"provider_type_id"`
	ResourceTypeID   int64     `gorm:"column:resource_type_id;primaryKey" json:"resource_type_id"`
	MetricKey        string    `gorm:"column:metric_key" json:"metric_key"`
	MetircDesc       string    `gorm:"column:metirc_desc" json:"metirc_desc"`
	MetricVvalueType string    `gorm:"column:metric_vvalue_type" json:"metric_vvalue_type"`
	Unit             string    `gorm:"column:unit" json:"unit"`
	DeleteYn         string    `gorm:"column:delete_yn" json:"delete_yn"`
	RegDt            time.Time `gorm:"column:reg_dt" json:"reg_dt"`
	ModDt            time.Time `gorm:"column:mod_dt" json:"mod_dt"`
	DelDt            time.Time `gorm:"column:del_dt" json:"del_dt"`
}

// TableName CmMetric's table name
func (*CmMetric) TableName() string {
	return TableNameCmMetric
}
