package routes

import (
	"GoBlog/api/v1"
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// User模块的路由接口
		router.POST("user/add", v1.AddUser)

		router.GET("users", v1.GetUsers)

		router.PUT("user/:id", v1.EditUser)

		router.DELETE("user/:id", v1.DeleteUser)

		// 分类的模块的路由接口

		// 文章模块的路由接口
	}

	r.Run(utils.HttpPort)
}
