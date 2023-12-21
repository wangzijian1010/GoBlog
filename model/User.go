package model

import (
	"GoBlog/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
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
	data.Password = ScryptPW(data.Password)
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
// 编辑用户密码不应该存在于这里面 应该限制在密码以外
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
// 现在都是软删除不会完全删除
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 密码加密 不使用明文保存
// 用钩子函数实现调用
//func (u *User) BeforeSave() {
//	u.Password = ScryptPW(u.Password)
//}

func ScryptPW(password string) string {
	const KeyLen = 10 // 生成的盐值
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	// 得到Hash数组 HashPw
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}

	finalPW := base64.StdEncoding.EncodeToString(HashPw)
	// 得到了加密后的密码如何写入数据库呢
	// 利用gorm的钩子函数在传入数据库之前进行更改
	return finalPW
}

func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USERNAME_NOT_EXIST
	}
	if ScryptPW(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
