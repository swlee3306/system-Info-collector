package db

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCbRack = "cb_rack"

// CbRack mapped from table <cb_rack>
type CbRack struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:렉 아이디" json:"id"` // 렉 아이디
	Name          string         `gorm:"column:name;comment:렉 이름" json:"name"`                            // 렉 이름
	CenterID      int64          `gorm:"column:center_id;not null;comment:렉 센터 아이디" json:"center_id"`     // 렉 센터 아이디
	Size          int32          `gorm:"column:size;comment:렉 내 자원의 수량" json:"size"`                      // 렉 내 자원의 수량
	AvailableUnit string         `gorm:"column:available_unit;comment:사용 가능 유닛" json:"available_unit"`    // 사용 가능 유닛
	UsingUnit     string         `gorm:"column:using_unit;comment:사용 중인 유닛" json:"using_unit"`            // 사용 중인 유닛
	DisableUnit   string         `gorm:"column:disable_unit;comment:비활성화 유닛" json:"disable_unit"`         // 비활성화 유닛
	Usable        string         `gorm:"column:usable;comment:용도" json:"usable"`                          // 용도
	Floor         string         `gorm:"column:floor;comment:층" json:"floor"`                             // 층
	Area          string         `gorm:"column:area;comment:구역" json:"area"`                              // 구역
	InstallDate   string         `gorm:"column:install_date;comment:설치일자" json:"install_date"`            // 설치일자
	Company       string         `gorm:"column:company;comment:제조사" json:"company"`                       // 제조사
	Model         string         `gorm:"column:model;comment:모델" json:"model"`                            // 모델
	Series        string         `gorm:"column:series;comment:시리즈" json:"series"`                         // 시리즈
	Source        string         `gorm:"column:source;comment:출처" json:"source"`                          // 출처
	Price         string         `gorm:"column:price;comment:가격" json:"price"`                            // 가격
	CompanyCharge string         `gorm:"column:company_charge;comment:담당업체" json:"company_charge"`        // 담당업체
	Manager       string         `gorm:"column:manager;comment:담당자" json:"manager"`                       // 담당자
	Contact       string         `gorm:"column:contact;comment:연락처" json:"contact"`                       // 연락처
	CreatedAt     time.Time      `gorm:"column:created_at;comment:생성 일시" json:"created_at"`               // 생성 일시
	UpdatedAt     time.Time      `gorm:"column:updated_at;comment:수정 일시" json:"updated_at"`               // 수정 일시
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;comment:삭제 일시" json:"deleted_at"`               // 삭제 일시
	Deleted       bool           `gorm:"column:deleted;comment:삭제 상태" json:"deleted"`                     // 삭제 상태
}

// TableName CbRack's table name
func (*CbRack) TableName() string {
	return TableNameCbRack
}
