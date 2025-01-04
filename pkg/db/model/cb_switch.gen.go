package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbSwitch = "cb_switch"

// CbSwitch mapped from table <cb_switch>
type CbSwitch struct {
	ID                int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:스위치 아이디" json:"id"`             // 스위치 아이디
	RackID            int64          `gorm:"column:rack_id;not null;comment:렉 아이디" json:"rack_id"`                          // 렉 아이디
	UUID              string         `gorm:"column:uuid;not null;default:uuid();comment:스위치 UUID" json:"uuid"`              // 스위치 UUID
	Name              string         `gorm:"column:name;comment:스위치 이름" json:"name"`                                        // 스위치 이름
	Status            string         `gorm:"column:status;comment:스위치 상태" json:"status"`                                    // 스위치 상태
	Location          string         `gorm:"column:location;comment:스위치 위치" json:"location"`                                // 스위치 위치
	Model             string         `gorm:"column:model;comment:스위치 모델 이름" json:"model"`                                   // 스위치 모델 이름
	NetworkUsage      string         `gorm:"column:network_usage;comment:스위치 네트워크 사용량" json:"network_usage"`                // 스위치 네트워크 사용량
	AssetRegistration string         `gorm:"column:asset_registration;comment:자산 등록" json:"asset_registration"`             // 자산 등록
	InstallDate       string         `gorm:"column:install_date;comment:설치 일자" json:"install_date"`                         // 설치 일자
	Company           string         `gorm:"column:company;comment:제조사" json:"company"`                                     // 제조사
	SerialNum         string         `gorm:"column:serial_num;comment:시리얼 번호" json:"serial_num"`                            // 시리얼 번호
	PortInfo          string         `gorm:"column:port_info;comment:포트구성" json:"port_info"`                                // 포트구성
	Usable            string         `gorm:"column:usable;comment:용도" json:"usable"`                                        // 용도
	GroupYn           string         `gorm:"column:group_yn;comment:그룹 여부" json:"group_yn"`                                 // 그룹 여부
	GroupPeerID       int64          `gorm:"column:group_peer_id;comment:그룹 상대 스위치 아이디" json:"group_peer_id"`               // 그룹 상대 스위치 아이디
	StackYn           string         `gorm:"column:stack_yn;comment:스태킹 여부" json:"stack_yn"`                                // 스태킹 여부
	StackType         string         `gorm:"column:stack_type;comment:스태킹 타입" json:"stack_type"`                            // 스태킹 타입
	SwitchPortFilter  string         `gorm:"column:switch_port_filter;comment:스위치 포트 필터" json:"switch_port_filter"`         // 스위치 포트 필터
	CreatedAt         time.Time      `gorm:"column:created_at;default:current_timestamp();comment:생성 일시" json:"created_at"` // 생성 일시
	UpdatedAt         time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                             // 수정 일시
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                             // 삭제 일시
	Deleted           bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                                   // 삭제 상태
}

// TableName CbSwitch's table name
func (*CbSwitch) TableName() string {
	return TableNameCbSwitch
}
