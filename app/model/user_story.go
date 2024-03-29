package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// UserStory Model
type UserStory struct {
	gorm.Model
	// Title       string          `json:"title"`
	Body string `json:"body"`
	// TitleVector   pq.Float64Array `gorm:"type:double precision[]" json:"-"`
	BodyVector    pq.Float64Array `gorm:"type:double precision[]" json:"-"`
	TagID         int             `json:"tagId"`
	Score         float64         `json:"Score"`
	Capability    string          `json:"capability"`
	SubCapability string          `json:"subcapability"`
	Epic          string          `json:"epic"`
	Result        string          `json:"result"`
}

//UserStoryExpand Model
type UserStoryExpand struct {
	gorm.Model
	Body       string          `json:"body"`
	Result     string          `json:"result"`
	BodyVector pq.Float64Array `gorm:"type:double precision[]" json:"-"`
	Score      float64         `json:"Score"`
}
