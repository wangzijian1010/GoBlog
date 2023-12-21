package v1

import (
	"GoBlog/middleware"
	"GoBlog/model"
	"GoBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password)

	// 如果成功了才生成token
	if code == errmsg.SUCCSE {
		token, code = middleware.SetToken(data.Username, data.Password)

	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
