package model

import (
	"GoBlog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CreateCate(data *Category) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询用户是否存在
func CheckCate(name string) (code int) {
	var cate Category
	// 这里是db的标准化输入 用来查询的 将查询到的结果输入到User结构体当中
	db.Where("name = ?", name).First(&cate)
	if cate.Name != "" {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetCate(pageSize int, PageNum int) []Category {
	var cate []Category
	// 进行分页查询 主要是为了如果SQl中的数据太多的情况下 无法在一个页面里面输出
	// 这里就是限制输出 并将查询到的记录写道users当中
	err = db.Limit(pageSize).Offset((PageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑用户
// 编辑用户密码不应该存在于这里面 应该限制在密码以外
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err := db.Model(&cate).Where("id = ?", id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
// 现在都是软删除不会完全删除
func DeleteCate(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// TODO 查询分类下的所有文章
