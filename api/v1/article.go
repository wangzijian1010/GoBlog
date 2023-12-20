package v1

import (
	"GoBlog/model"
	"GoBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddArticle(c *gin.Context) {
	// 定义一个User的结构体
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	// 利用前端传入的数值绑定到data上面
	code := model.CreateArt(&data)

	// 返回JSON 调用errmsg的GetMSg方法
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// TODO 查询分类下的所有文章

// TODO 利用文章名来查询单个文章

// 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) // 这个返回的字符串 要将id转换为int

	code := model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 编辑文章
func EditArticle(c *gin.Context) {
	// 首先先要找到对应的用户名的用户数据
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Article
	// 更新的时候也可以看作为post他会提交表单
	c.ShouldBindJSON(&data)

	code := model.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}
