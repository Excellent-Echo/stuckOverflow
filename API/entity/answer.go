package entity

import (
	"time"

	"gorm.io/gorm"
)

type Answers struct {
	ID         uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Content    string         `gorm:"text;not null" json:"content"`
	ImageFile  string         `gorm:"text" json:"image_file"`
	UserID     uint32         `json:"user_id"`
	QuestionID uint64         `json:"question_id"`
	CreatedAt  time.Time      `gorm:"type:datetime;not null;default:current_timestamp" json:"-"`
	UpdatedAt  time.Time      `gorm:"type:datetime;default:current_timestamp" json:"-"`
	Deleted    gorm.DeletedAt `json:"-"`
}
