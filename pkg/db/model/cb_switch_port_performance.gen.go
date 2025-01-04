package db

const TableNameCbSwitchPortPerformance = "cb_switch_port_performance"

// CbSwitchPortPerformance mapped from table <cb_switch_port_performance>
type CbSwitchPortPerformance struct {
	SwitchPortID            int64  `gorm:"column:switch_port_id;primaryKey;comment:스위치 포트 아이디" json:"switch_port_id"`                            // 스위치 포트 아이디
	SwitchPortBandwidth     string `gorm:"column:switch_port_bandwidth;comment:포트별 대역폭" json:"switch_port_bandwidth"`                            // 포트별 대역폭
	SwitchPortInBps         string `gorm:"column:switch_port_in_bps;comment:포트별 초당 사용량(in)" json:"switch_port_in_bps"`                           // 포트별 초당 사용량(in)
	SwitchPortOutBps        string `gorm:"column:switch_port_out_bps;comment:포트별 초당 사용량(out)" json:"switch_port_out_bps"`                        // 포트별 초당 사용량(out)
	SwitchPortInUsage       string `gorm:"column:switch_port_in_usage;comment:포트별 초당 사용률(in)" json:"switch_port_in_usage"`                       // 포트별 초당 사용률(in)
	SwitchPortOutUsage      string `gorm:"column:switch_port_out_usage;comment:포트별 초당 사용률(out)" json:"switch_port_out_usage"`                    // 포트별 초당 사용률(out)
	SwitchPortInUcastUsage  string `gorm:"column:switch_port_in_ucast_usage;comment:포트별 초당 사용률(유니캐스트 in)" json:"switch_port_in_ucast_usage"`     // 포트별 초당 사용률(유니캐스트 in)
	SwitchPortInMcastUsage  string `gorm:"column:switch_port_in_mcast_usage;comment:포트별 초당 사용률(멀티캐스트 in)" json:"switch_port_in_mcast_usage"`     // 포트별 초당 사용률(멀티캐스트 in)
	SwitchPortInBcastUsage  string `gorm:"column:switch_port_in_bcast_usage;comment:포트별 초당 사용률(브로드캐스트 in)" json:"switch_port_in_bcast_usage"`    // 포트별 초당 사용률(브로드캐스트 in)
	SwitchPortOutUcastUsage string `gorm:"column:switch_port_out_ucast_usage;comment:포트별 초당 사용률(유니캐스트 out)" json:"switch_port_out_ucast_usage"`  // 포트별 초당 사용률(유니캐스트 out)
	SwitchPortOutMcastUsage string `gorm:"column:switch_port_out_mcast_usage;comment:포트별 초당 사용률(멀티캐스트 out)" json:"switch_port_out_mcast_usage"`  // 포트별 초당 사용률(멀티캐스트 out)
	SwitchPortOutBcastUsage string `gorm:"column:switch_port_out_bcast_usage;comment:포트별 초당 사용률(브로드캐스트 out)" json:"switch_port_out_bcast_usage"` // 포트별 초당 사용률(브로드캐스트 out)
}

// TableName CbSwitchPortPerformance's table name
func (*CbSwitchPortPerformance) TableName() string {
	return TableNameCbSwitchPortPerformance
}
