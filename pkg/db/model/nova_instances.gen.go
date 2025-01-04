package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameNovaInstance = "nova_instances"

// NovaInstance mapped from table <nova_instances>
type NovaInstance struct {
	ProviderID             string         `gorm:"column:provider_id;not null" json:"provider_id"`
	Domain                 string         `gorm:"column:domain" json:"domain"`
	CreatedAt              time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ID                     int32          `gorm:"column:id;primaryKey" json:"id"`
	ProjectID              string         `gorm:"column:project_id" json:"project_id"`
	UserID                 string         `gorm:"column:user_id" json:"user_id"`
	ImageRef               string         `gorm:"column:image_ref" json:"image_ref"`
	KernelID               string         `gorm:"column:kernel_id" json:"kernel_id"`
	RamdiskID              string         `gorm:"column:ramdisk_id" json:"ramdisk_id"`
	LaunchIndex            int32          `gorm:"column:launch_index" json:"launch_index"`
	KeyName                string         `gorm:"column:key_name" json:"key_name"`
	KeyData                string         `gorm:"column:key_data" json:"key_data"`
	PowerState             int32          `gorm:"column:power_state" json:"power_state"`
	VMState                string         `gorm:"column:vm_state" json:"vm_state"`
	MemoryMb               int32          `gorm:"column:memory_mb" json:"memory_mb"`
	Vcpus                  int32          `gorm:"column:vcpus" json:"vcpus"`
	VMName                 string         `gorm:"column:vm_name" json:"vm_name"`
	Host                   string         `gorm:"column:host" json:"host"`
	UserData               string         `gorm:"column:user_data" json:"user_data"`
	ReservationID          string         `gorm:"column:reservation_id" json:"reservation_id"`
	LaunchedAt             time.Time      `gorm:"column:launched_at" json:"launched_at"`
	TerminatedAt           time.Time      `gorm:"column:terminated_at" json:"terminated_at"`
	DisplayName            string         `gorm:"column:display_name" json:"display_name"`
	DisplayDescription     string         `gorm:"column:display_description" json:"display_description"`
	AvailabilityZone       string         `gorm:"column:availability_zone" json:"availability_zone"`
	Locked                 bool           `gorm:"column:locked" json:"locked"`
	OsType                 string         `gorm:"column:os_type" json:"os_type"`
	LaunchedOn             string         `gorm:"column:launched_on" json:"launched_on"`
	InstanceTypeID         int32          `gorm:"column:instance_type_id" json:"instance_type_id"`
	VMMode                 string         `gorm:"column:vm_mode" json:"vm_mode"`
	UUID                   string         `gorm:"column:uuid;not null" json:"uuid"`
	Architecture           string         `gorm:"column:architecture" json:"architecture"`
	RootDeviceName         string         `gorm:"column:root_device_name" json:"root_device_name"`
	AccessIPV4             string         `gorm:"column:access_ip_v4" json:"access_ip_v4"`
	AccessIPV6             string         `gorm:"column:access_ip_v6" json:"access_ip_v6"`
	ConfigDrive            string         `gorm:"column:config_drive" json:"config_drive"`
	TaskState              string         `gorm:"column:task_state" json:"task_state"`
	DefaultEphemeralDevice string         `gorm:"column:default_ephemeral_device" json:"default_ephemeral_device"`
	DefaultSwapDevice      string         `gorm:"column:default_swap_device" json:"default_swap_device"`
	Progress               int32          `gorm:"column:progress" json:"progress"`
	AutoDiskConfig         bool           `gorm:"column:auto_disk_config" json:"auto_disk_config"`
	ShutdownTerminate      bool           `gorm:"column:shutdown_terminate" json:"shutdown_terminate"`
	DisableTerminate       bool           `gorm:"column:disable_terminate" json:"disable_terminate"`
	RootGb                 int32          `gorm:"column:root_gb" json:"root_gb"`
	EphemeralGb            int32          `gorm:"column:ephemeral_gb" json:"ephemeral_gb"`
	CellName               string         `gorm:"column:cell_name" json:"cell_name"`
	Node                   string         `gorm:"column:node" json:"node"`
	Deleted                int32          `gorm:"column:deleted" json:"deleted"`
	LockedBy               string         `gorm:"column:locked_by" json:"locked_by"`
	Cleaned                int32          `gorm:"column:cleaned" json:"cleaned"`
	EphemeralKeyUUID       string         `gorm:"column:ephemeral_key_uuid" json:"ephemeral_key_uuid"`
	Hidden                 bool           `gorm:"column:hidden" json:"hidden"`
}

// TableName NovaInstance's table name
func (*NovaInstance) TableName() string {
	return TableNameNovaInstance
}
