package v1

import (
	"GoBlog/model"
	"GoBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCate(c *gin.Context) {
	// 定义一个User的结构体
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	// 利用前端传入的数值绑定到data上面
	// 调用model的CheckUser函数来确定里面是否有同名的名称 并返回code
	code := model.CheckCate(data.Name)
	// 利用code来判断是否有重名函数
	if code == errmsg.SUCCSE {
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}

	// 返回JSON 调用errmsg的GetMSg方法
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	// 如果pagesize = 0 那就可以不分页
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetCate(pageSize, pageNum)
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
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) // 这个返回的字符串 要将id转换为int

	code := model.DeleteCate(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 编辑分类名
func EditCate(c *gin.Context) {
	// 首先先要找到对应的用户名的用户数据
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	// 更新的时候也可以看作为post他会提交表单
	c.ShouldBindJSON(&data)
	// 检查分类名是否存在
	code := model.CheckCate(data.Name)

	if code == errmsg.SUCCSE {
		model.EditCate(id, &data)
	}

	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
