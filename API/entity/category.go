package entity

import (
	"time"
)

type Categories struct {
	ID           uint32      `gorm:"primaryKey;autoIncrement;" json:"id"`
	CategoryName string      `gorm:"size:255; not null;unique" json:"category_name"`
	Description  string      `json:"description"`
	CreatedAt    time.Time   `gorm:"type:datetime;not null;default:current_timestamp" json:"-"`
	UpdatedAt    time.Time   `gorm:"type:datetime;not null" json:"-"`
	Questions    []Questions `gorm:"foreignKey:CategoryID" json:"questions"`
}
