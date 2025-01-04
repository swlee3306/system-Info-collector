package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbHost = "cb_host"

// CbHost mapped from table <cb_host>
type CbHost struct {
	ID                 int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:호스트 아이디" json:"id"`             // 호스트 아이디
	RackID             int64          `gorm:"column:rack_id;not null;comment:렉 아이디" json:"rack_id"`                          // 렉 아이디
	UUID               string         `gorm:"column:uuid;not null;default:uuid();comment:호스트 UUID" json:"uuid"`              // 호스트 UUID
	Name               string         `gorm:"column:name;comment:호스트 이름" json:"name"`                                        // 호스트 이름
	Status             string         `gorm:"column:status;comment:호스트 상태" json:"status"`                                    // 호스트 상태
	HostGroupID        string         `gorm:"column:hostGroupId;comment:호스트 그룹 아이디" json:"hostGroupId"`                      // 호스트 그룹 아이디
	AssetRegistration  string         `gorm:"column:asset_registration;comment:자산 등록" json:"asset_registration"`             // 자산 등록
	InstallDate        string         `gorm:"column:install_date;comment:설치 일자" json:"install_date"`                         // 설치 일자
	Company            string         `gorm:"column:company;comment:제조사" json:"company"`                                     // 제조사
	SerialNum          string         `gorm:"column:serial_num;comment:시리얼 번호" json:"serial_num"`                            // 시리얼 번호
	Model              string         `gorm:"column:model;comment:모델 명" json:"model"`                                        // 모델 명
	OsInfo             string         `gorm:"column:os_info;comment:운영 체제" json:"os_info"`                                   // 운영 체제
	ConsoleInfo        string         `gorm:"column:console_info;comment:관리 콘솔 정보" json:"console_info"`                      // 관리 콘솔 정보
	Location           string         `gorm:"column:location;comment:위치" json:"location"`                                    // 위치
	IpmiIP             string         `gorm:"column:ipmi_ip;comment:IPMI용 아이피" json:"ipmi_ip"`                               // IPMI용 아이피
	IpmiPort           int32          `gorm:"column:ipmi_port;comment:IPMI용 포트" json:"ipmi_port"`                            // IPMI용 포트
	IpmiUserid         string         `gorm:"column:ipmi_userid;comment:IPMI용 아이디" json:"ipmi_userid"`                       // IPMI용 아이디
	IpmiUserpwd        string         `gorm:"column:ipmi_userpwd;comment:IPMI용 비밀번호" json:"ipmi_userpwd"`                    // IPMI용 비밀번호
	SnmpIP             string         `gorm:"column:snmp_ip;comment:수집용 snmp 아이피" json:"snmp_ip"`                            // 수집용 snmp 아이피
	SnmpPort           int32          `gorm:"column:snmp_port;comment:수집용 snmp 포트" json:"snmp_port"`                         // 수집용 snmp 포트
	SnmpVersion        string         `gorm:"column:snmp_version;comment:수집용 snmp 버전" json:"snmp_version"`                   // 수집용 snmp 버전
	SnmpCommunity      string         `gorm:"column:snmp_community;comment:수집용 snmp 커뮤니티명" json:"snmp_community"`            // 수집용 snmp 커뮤니티명
	HypervisorHostID   string         `gorm:"column:hypervisor_host_id;comment:하이퍼바이저 호스트 ID" json:"hypervisor_host_id"`     // 하이퍼바이저 호스트 ID
	HypervisorHostName string         `gorm:"column:hypervisor_host_name;comment:하이퍼바이저 호스트 이름" json:"hypervisor_host_name"` // 하이퍼바이저 호스트 이름
	CreatedAt          time.Time      `gorm:"column:created_at;default:current_timestamp();comment:생성 일시" json:"created_at"` // 생성 일시
	UpdatedAt          time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`                             // 수정 일시
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`                             // 삭제 일시
	Deleted            bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                                   // 삭제 상태
	Usable             string         `gorm:"column:usable;comment:용도" json:"usable"`                                        // 용도
	VMCount            int32          `gorm:"column:vm_count;comment:vm 갯" json:"vm_count"`                                  // vm 갯
}

// TableName CbHost's table name
func (*CbHost) TableName() string {
	return TableNameCbHost
}
