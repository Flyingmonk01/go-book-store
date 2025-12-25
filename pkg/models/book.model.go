package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	ID            string `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"type:varchar(100);not null" json:"name"`
	Author        string `gorm:"type:varchar(100);not null" json:"author"`
	PublishedYear int    `gorm:"not null" json:"published_year"`
}
