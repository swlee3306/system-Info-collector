package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameNovaInstanceInfoCach = "nova_instance_info_caches"

// NovaInstanceInfoCach mapped from table <nova_instance_info_caches>
type NovaInstanceInfoCach struct {
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID           int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NetworkInfo  string         `gorm:"column:network_info" json:"network_info"`
	InstanceUUID string         `gorm:"column:instance_uuid;not null" json:"instance_uuid"`
	Deleted      int32          `gorm:"column:deleted" json:"deleted"`
}

// TableName NovaInstanceInfoCach's table name
func (*NovaInstanceInfoCach) TableName() string {
	return TableNameNovaInstanceInfoCach
}
