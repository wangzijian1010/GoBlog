package v1

import (
	"GoBlog/model"
	"GoBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {
	// 想查询就直接调用model.CheckUser
	var data model.User
	_ = c.ShouldBindJSON(&data)
	// 由于姓名是无法重复的 所以可以按照姓名来查询用户
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_NOT_EXIST {
		code = errmsg.ERROR_USERNAME_NOT_EXIST
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 添加用户
func AddUser(c *gin.Context) {
	// 定义一个User的结构体
	var data model.User
	_ = c.ShouldBindJSON(&data)
	// 利用前端传入的数值绑定到data上面
	// 调用model的CheckUser函数来确定里面是否有同名的名称 并返回code
	code := model.CheckUser(data.Username)
	// 利用code来判断是否有重名函数
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	// 返回JSON 调用errmsg的GetMSg方法
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	// 如果pagesize = 0 那就可以不分页
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除用户
func DeleteUser(c *gin.Context) {

}

// 编辑用户
func EditUser(c *gin.Context) {

}

// 查询单个用户