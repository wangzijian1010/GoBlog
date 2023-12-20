package model

import (
	"GoBlog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	// 这里的Json : title 就是对外暴露的是title形式 简单粗暴理解就是在json中你输入title就对应这里的Title
	Title      string   `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID uint     `gorm:"not null"`              // 这是外键
	Category   Category `gorm:"foreignKey:CategoryID"` // 关联Category
	Desc       string   `gorm:"type:varchar(200)" json:"desc"`
	Content    string   `gorm:"type:longtext" json:"content"`
	Img        string   `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArt(data *Article) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// TODO 查询分类下的所有文章

// TODO 查询单个文章

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.CategoryID
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err := db.Model(&art).Where("id = ?", id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
