package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbSwitchPort = "cb_switch_port"

// CbSwitchPort mapped from table <cb_switch_port>
type CbSwitchPort struct {
	ID                       int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:스위치 포트 아이디" json:"id"` // 스위치 포트 아이디
	SwitchesID               int64          `gorm:"column:switches_id;not null;comment:스위치 아이디" json:"switches_id"`       // 스위치 아이디
	SwitchesIDPhysical       int64          `gorm:"column:switches_id_physical" json:"switches_id_physical"`
	PortNo                   int64          `gorm:"column:port_no" json:"port_no"`
	PortName                 string         `gorm:"column:port_name" json:"port_name"`
	Status                   string         `gorm:"column:status;comment:스위치 상태" json:"status"`         // 스위치 상태
	MacAddress               string         `gorm:"column:mac_address;comment:맥주소" json:"mac_address"`  // 맥주소
	IPAddress                string         `gorm:"column:ip_address;comment:아이피 주소" json:"ip_address"` // 아이피 주소
	PeerID                   int64          `gorm:"column:peer_id;comment:연결 아이디" json:"peer_id"`       // 연결 아이디
	PeerType                 string         `gorm:"column:peer_type;comment:연결 타입" json:"peer_type"`    // 연결 타입
	PortIfType               string         `gorm:"column:port_ifType" json:"port_ifType"`
	PortIfMtu                string         `gorm:"column:port_ifMtu" json:"port_ifMtu"`
	PortIfHighSpeed          string         `gorm:"column:port_ifHighSpeed" json:"port_ifHighSpeed"`
	PortIfPhysAddress        string         `gorm:"column:port_ifPhysAddress" json:"port_ifPhysAddress"`
	PortIfAdminStatus        string         `gorm:"column:port_ifAdminStatus" json:"port_ifAdminStatus"`
	PortIfLastChange         time.Time      `gorm:"column:port_ifLastChange" json:"port_ifLastChange"`
	PortIfHCInOctets         string         `gorm:"column:port_ifHCInOctets" json:"port_ifHCInOctets"`
	PortIfHCInUcastPkts      string         `gorm:"column:port_ifHCInUcastPkts" json:"port_ifHCInUcastPkts"`
	PortIfHCInMulticastPkts  string         `gorm:"column:port_ifHCInMulticastPkts" json:"port_ifHCInMulticastPkts"`
	PortIfHCInBroadcastPkts  string         `gorm:"column:port_ifHCInBroadcastPkts" json:"port_ifHCInBroadcastPkts"`
	PortIfInDiscards         string         `gorm:"column:port_ifInDiscards" json:"port_ifInDiscards"`
	PortIfInErrors           string         `gorm:"column:port_ifInErrors" json:"port_ifInErrors"`
	PortIfInUnknownProtos    string         `gorm:"column:port_ifInUnknownProtos" json:"port_ifInUnknownProtos"`
	PortIfHCOutOctets        string         `gorm:"column:port_ifHCOutOctets" json:"port_ifHCOutOctets"`
	PortIfHCOutUcastPkts     string         `gorm:"column:port_ifHCOutUcastPkts" json:"port_ifHCOutUcastPkts"`
	PortIfHCOutMulticastPkts string         `gorm:"column:port_ifHCOutMulticastPkts" json:"port_ifHCOutMulticastPkts"`
	PortIfHCOutBroadcastPkts string         `gorm:"column:port_ifHCOutBroadcastPkts" json:"port_ifHCOutBroadcastPkts"`
	PortIfOutDiscards        string         `gorm:"column:port_ifOutDiscards" json:"port_ifOutDiscards"`
	PortIfOutErrors          string         `gorm:"column:port_ifOutErrors" json:"port_ifOutErrors"`
	PortIfOutQLen            string         `gorm:"column:port_ifOutQLen" json:"port_ifOutQLen"`
	PortIfSpecific           string         `gorm:"column:port_ifSpecific" json:"port_ifSpecific"`
	CreatedAt                time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"` // 생성 일시
	UpdatedAt                time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"` // 수정 일시
	DeletedAt                gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"` // 삭제 일시
	Deleted                  bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`       // 삭제 상태
}

// TableName CbSwitchPort's table name
func (*CbSwitchPort) TableName() string {
	return TableNameCbSwitchPort
}
