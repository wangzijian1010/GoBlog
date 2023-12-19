package model

import (
	"GoBlog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` // 利用role来确定权限
}

// 新增用户 这里面是对数据库的操作
// 就是v1中的接口会调用model中的操作来实现控制器的功能
// 这里的code就是状态码
func CreateUser(data *User) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户是否存在
func CheckUser(username string) (code int) {
	var users User
	// 这里是db的标准化输入 用来查询的 将查询到的结果输入到User结构体当中
	db.Where("username = ?", username).First(&users)
	if users.Username != "" {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetUsers(pageSize int, PageNum int) []User {
	var users []User
	// 进行分页查询 主要是为了如果SQl中的数据太多的情况下 无法在一个页面里面输出
	// 这里就是限制输出 并将查询到的记录写道users当中
	err = db.Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户

// 删除用户
