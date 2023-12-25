package v1

import (
	"GoBlog/model"
	"GoBlog/utils/errmsg"
	"GoBlog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询用户是否存在

// 添加用户
func AddUser(c *gin.Context) {
	// 定义一个User的结构体
	var data model.User
	_ = c.ShouldBindJSON(&data)
	// 利用前端传入的数值绑定到data上面

	var msg string
	var code int
	msg, code = validator.Validate(&data)

	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	// 调用model的CheckUser函数来确定里面是否有同名的名称 并返回code
	code = model.CheckUser(data.Username)
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

	data, total := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除用户
// 现在一般都是软删除
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) // 这个返回的字符串 要将id转换为int

	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 编辑用户
func EditUser(c *gin.Context) {
	// 首先先要找到对应的用户名的用户数据
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.User
	// 更新的时候也可以看作为post他会提交表单
	c.ShouldBindJSON(&data)

	code := model.CheckUser(data.Username)

	if code == errmsg.SUCCSE {
		model.EditUser(id, &data)
	}

	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
