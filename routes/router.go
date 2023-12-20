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
		router.POST("category/add", v1.AddCate)

		router.GET("category", v1.GetCate)

		router.PUT("category/:id", v1.EditCate)

		router.DELETE("category/:id", v1.DeleteCate)

		// 文章模块的路由接口
		router.POST("article/add", v1.AddArticle)

		// TODO router.GET("category", v1.GetArticle)

		router.PUT("article/:id", v1.EditArticle)

		router.DELETE("article/:id", v1.DeleteArticle)
	}

	r.Run(utils.HttpPort)
}
