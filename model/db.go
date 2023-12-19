package model

import (
	"GoBlog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	// 连接数据库
	//dsn := "root:123456@tcp(127.0.0.1:8888)/user?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := utils.DbUser + ":" + utils.DbPassWord + "@tcp(" + utils.DbHost + ":" + utils.DbPort + ")/" + utils.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("连接数据库出错", err)
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{})

}
