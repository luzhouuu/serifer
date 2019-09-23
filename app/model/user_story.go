package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// UserStory Model
type UserStory struct {
	gorm.Model
	// Title       string          `json:"title"`
	Body        string          `json:"body"`
	TitleVector pq.Float64Array `gorm:"type:double precision[]" json:"-"`
	BodyVector  pq.Float64Array `gorm:"type:double precision[]" json:"-"`
	TagID       uint            `json:"tagId"`
	Score       float64         `json:"Score"`
}
