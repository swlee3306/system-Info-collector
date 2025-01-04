package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameNovaBlockDeviceMapping = "nova_block_device_mapping"

// NovaBlockDeviceMapping mapped from table <nova_block_device_mapping>
type NovaBlockDeviceMapping struct {
	CreatedAt           time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID                  int32          `gorm:"column:id;primaryKey" json:"id"`
	DeviceName          string         `gorm:"column:device_name" json:"device_name"`
	DeleteOnTermination bool           `gorm:"column:delete_on_termination" json:"delete_on_termination"`
	SnapshotID          string         `gorm:"column:snapshot_id" json:"snapshot_id"`
	VolumeID            string         `gorm:"column:volume_id" json:"volume_id"`
	VolumeSize          int32          `gorm:"column:volume_size" json:"volume_size"`
	NoDevice            bool           `gorm:"column:no_device" json:"no_device"`
	ConnectionInfo      string         `gorm:"column:connection_info" json:"connection_info"`
	InstanceUUID        string         `gorm:"column:instance_uuid" json:"instance_uuid"`
	Deleted             int32          `gorm:"column:deleted" json:"deleted"`
	SourceType          string         `gorm:"column:source_type" json:"source_type"`
	DestinationType     string         `gorm:"column:destination_type" json:"destination_type"`
	GuestFormat         string         `gorm:"column:guest_format" json:"guest_format"`
	DeviceType          string         `gorm:"column:device_type" json:"device_type"`
	DiskBus             string         `gorm:"column:disk_bus" json:"disk_bus"`
	BootIndex           int32          `gorm:"column:boot_index" json:"boot_index"`
	ImageID             string         `gorm:"column:image_id" json:"image_id"`
	Tag                 string         `gorm:"column:tag" json:"tag"`
	AttachmentID        string         `gorm:"column:attachment_id" json:"attachment_id"`
	UUID                string         `gorm:"column:uuid" json:"uuid"`
	VolumeType          string         `gorm:"column:volume_type" json:"volume_type"`
}

// TableName NovaBlockDeviceMapping's table name
func (*NovaBlockDeviceMapping) TableName() string {
	return TableNameNovaBlockDeviceMapping
}
