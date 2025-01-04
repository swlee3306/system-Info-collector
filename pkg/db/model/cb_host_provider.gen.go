package db

const TableNameCbHostProvider = "cb_host_provider"

// CbHostProvider mapped from table <cb_host_provider>
type CbHostProvider struct {
	HostID       int64  `gorm:"column:host_id;comment:호스트 아이디" json:"host_id"`                   // 호스트 아이디
	HostName     string `gorm:"column:host_name;comment:호스트 이름" json:"host_name"`                // 호스트 이름
	ProviderID   string `gorm:"column:provider_id;comment:프로바이더 아이디" json:"provider_id"`         // 프로바이더 아이디
	ProviderUUID string `gorm:"column:provider_uuid;comment:provider uuid" json:"provider_uuid"` // provider uuid
}

// TableName CbHostProvider's table name
func (*CbHostProvider) TableName() string {
	return TableNameCbHostProvider
}
