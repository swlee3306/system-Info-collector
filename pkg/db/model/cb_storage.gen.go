package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbStorage = "cb_storage"

// CbStorage mapped from table <cb_storage>
type CbStorage struct {
	ID                int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:스토리지 아이디" json:"id"`            // 스토리지 아이디
	RackID            int64          `gorm:"column:rack_id;not null;comment:렉 아이디" json:"rack_id"`                          // 렉 아이디
	UUID              string         `gorm:"column:uuid;not null;default:uuid();comment:스토리지 UUID" json:"uuid"`             // 스토리지 UUID
	Status            string         `gorm:"column:status;comment:스토리지 상태" json:"status"`                                   // 스토리지 상태
	Name              string         `gorm:"column:name;comment:스토리지 명" json:"name"`                                        // 스토리지 명
	AssetRegistration string         `gorm:"column:asset_registration;comment:자산 등록" json:"asset_registration"`             // 자산 등록
	InstallDate       string         `gorm:"column:install_date;comment:설치 일자" json:"install_date"`                         // 설치 일자
	Company           string         `gorm:"column:company;comment:제조사" json:"company"`                                     // 제조사
	SerialNum         string         `gorm:"column:serial_num;comment:시리얼 번호" json:"serial_num"`                            // 시리얼 번호
	Model             string         `gorm:"column:model;comment:모델 명" json:"model"`                                        // 모델 명
	SnmpIP            string         `gorm:"column:snmp_ip;comment:수집용 snmp 아이피" json:"snmp_ip"`                            // 수집용 snmp 아이피
	SnmpPort          int32          `gorm:"column:snmp_port;comment:수집용 snmp 포트" json:"snmp_port"`                         // 수집용 snmp 포트
	SnmpVersion       string         `gorm:"column:snmp_version;comment:수집용 snmp 버전" json:"snmp_version"`                   // 수집용 snmp 버전
	SnmpCommunity     string         `gorm:"column:snmp_community;comment:수집용 snmp 커뮤니티명" json:"snmp_community"`            // 수집용 snmp 커뮤니티명
	CreatedAt         time.Time      `gorm:"column:created_at;default:current_timestamp();comment:생성 일시" json:"created_at"` // 생성 일시
	UpdatedAt         time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                             // 수정 일시
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                             // 삭제 일시
	Deleted           bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                                   // 삭제 상태
	VMCount           int32          `gorm:"column:vm_count;comment:vm 갯수" json:"vm_count"`                                 // vm 갯수
	Usable            string         `gorm:"column:usable;comment:용도" json:"usable"`                                        // 용도
}

// TableName CbStorage's table name
func (*CbStorage) TableName() string {
	return TableNameCbStorage
}
