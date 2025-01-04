package db

import (
	"time"
)

const TableNameGlanceImageTag = "glance_image_tags"

// GlanceImageTag mapped from table <glance_image_tags>
type GlanceImageTag struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	Value     string    `gorm:"column:value;not null" json:"value"`
	ImageID   string    `gorm:"column:image_id" json:"image_id"`
	Deleted   bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName GlanceImageTag's table name
func (*GlanceImageTag) TableName() string {
	return TableNameGlanceImageTag
}
