package db

import (
	"time"
)

const TableNameGlanceImage = "glance_images"

// GlanceImage mapped from table <glance_images>
type GlanceImage struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	Deleted   bool      `gorm:"column:deleted;not null" json:"deleted"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	OwnerID   string    `gorm:"column:owner_id;not null" json:"owner_id"`
	Size      int64     `gorm:"column:size;not null" json:"size"`
	Filename  string    `gorm:"column:filename" json:"filename"`
	Savename  string    `gorm:"column:savename" json:"savename"`
	Isimage   bool      `gorm:"column:isimage" json:"isimage"`
}

// TableName GlanceImage's table name
func (*GlanceImage) TableName() string {
	return TableNameGlanceImage
}
