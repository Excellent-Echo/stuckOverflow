package entity

import (
	"time"

	"gorm.io/gorm"
)

type Questions struct {
	ID         uint64         `gorm:"primaryKey;autoIncrement;" json:"id"`
	Title      string         `gorm:"size:255;not null" json:"title"`
	Content    string         `gorm:"text;not null" json:"content"`
	ImageFile  string         `gorm:"text" json:"image_file"`
	UserID     uint32         `json:"user_id"`
	User       User           `gorm:"foreignKey:UserID" json:"user_detail"`
	CategoryID uint32         `json:"category_id"`
	Category   Categories     `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	Answers    []Answers      `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"answers"`
	CreatedAt  time.Time      `gorm:"type:datetime" json:"-"`
	UpdatedAt  time.Time      `gorm:"type:datetime" json:"-"`
	Deleted    gorm.DeletedAt `json:"-"`
}
