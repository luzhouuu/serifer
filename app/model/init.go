package model

import "github.com/jinzhu/gorm"

// DB pointer
var DB *gorm.DB

// Init Model
func Init(db *gorm.DB) {
	DB = db
	DB.AutoMigrate(&UserStory{})
}

//InitExpand Model
func InitExpand(db *gorm.DB) {
	DB = db
	DB.AutoMigrate(&UserStoryExpand{})
}
