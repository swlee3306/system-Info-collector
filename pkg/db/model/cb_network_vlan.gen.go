package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbNetworkVlan = "cb_network_vlan"

// CbNetworkVlan mapped from table <cb_network_vlan>
type CbNetworkVlan struct {
	ID                      int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProviderID             int64          `gorm:"column:provider_id;not null"`
	Name                    string         `gorm:"column:name;not null" json:"name"`
	Status                  string         `gorm:"column:status;comment:ACTIVE, DOWN, BUILD, ERROR" json:"status"` // ACTIVE, DOWN, BUILD, ERROR
	UUID                    string         `gorm:"column:uuid;not null" json:"uuid"`
	TenantID                string         `gorm:"column:tenant_id" json:"tenant_id"`
	ProjectID               string         `gorm:"column:project_id" json:"project_id"`
	ProviderNetworkType     string         `gorm:"column:provider_network_type" json:"provider_network_type"`
	ProviderPhysicalNetwork string         `gorm:"column:provider_physical_network" json:"provider_physical_network"`
	ProviderSegmentationID  int32          `gorm:"column:provider_segmentation_id" json:"provider_segmentation_id"`
	Subnets                 string         `gorm:"column:subnets" json:"subnets"`
	SubnetsCidr             string         `gorm:"column:subnets_cidr"`
	Tags                    string         `gorm:"column:tags" json:"tags"`
	RevisionNumber          int32          `gorm:"column:revision_number" json:"revision_number"`
	CreatedAt               time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"`    // 생성 일시
	UpdatedAt               time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`    // 수정 일시
	DeletedAt               gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`    // 삭제 일시
	Deleted                 bool           `gorm:"column:deleted;not null;comment:삭제 상태" json:"deleted"` // 삭제 상태
	// 2024.11.04: bon.add
	IVal struct {
		IsExist bool // fase 일 경우, db 에 delete 된다.
	} `gorm:"-"`
}

// TableName CbNetworkVlan's table name
func (*CbNetworkVlan) TableName() string {
	return TableNameCbNetworkVlan
}