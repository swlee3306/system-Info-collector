package db

import (
	"time"
)

const TableNameNovaAPIFlavor = "nova_api_flavors"

// NovaAPIFlavor mapped from table <nova_api_flavors>
type NovaAPIFlavor struct {
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	ID          int32     `gorm:"column:id;primaryKey" json:"id"`
	MemoryMb    int32     `gorm:"column:memory_mb;not null" json:"memory_mb"`
	Vcpus       int32     `gorm:"column:vcpus;not null" json:"vcpus"`
	Swap        int32     `gorm:"column:swap;not null" json:"swap"`
	VcpuWeight  int32     `gorm:"column:vcpu_weight" json:"vcpu_weight"`
	Flavorid    string    `gorm:"column:flavorid;not null" json:"flavorid"`
	RxtxFactor  float32   `gorm:"column:rxtx_factor" json:"rxtx_factor"`
	RootGb      int32     `gorm:"column:root_gb" json:"root_gb"`
	EphemeralGb int32     `gorm:"column:ephemeral_gb" json:"ephemeral_gb"`
	Disabled    bool      `gorm:"column:disabled" json:"disabled"`
	IsPublic    bool      `gorm:"column:is_public" json:"is_public"`
	Description string    `gorm:"column:description" json:"description"`
}

// TableName NovaAPIFlavor's table name
func (*NovaAPIFlavor) TableName() string {
	return TableNameNovaAPIFlavor
}
