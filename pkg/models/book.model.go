package models

import (
	"github.com/Flyingmonk01/go-book-store/pkg/config"
	"github.com/jinzhu/gorm"
)

var db = config.GetDB()

type Book struct {
	gorm.Model
	ID            string `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"type:varchar(100);not null" json:"name"`
	Author        string `gorm:"type:varchar(100);not null" json:"author"`
	PublishedYear int    `gorm:"not null" json:"published_year"`
}

func init() {
	config.GetDB()
	db = config.GetDB()

	// db.AutoMigrate(&Book{})
}
