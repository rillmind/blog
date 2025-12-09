package posts

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	Title   string `json:"title" gorm:"type:varchar(50)"`
	Content string `json:"content" gorm:"type:text"`
	Slug    string `json:"slug"`
}

type ModelResponse struct {
	ID        uint
	Title     string
	Slug      string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
