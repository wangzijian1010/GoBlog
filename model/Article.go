package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title      string   `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID uint     `gorm:"not null"`              // 这是外键
	Category   Category `gorm:"foreignKey:CategoryID"` // 关联Category
	Desc       string   `gorm:"type:varchar(200)" json:"desc"`
	Content    string   `gorm:"type:longtext" json:"content"`
	Img        string   `gorm:"type:varchar(100)" json:"img"`
}
