package routes

import (
	"GoBlog/api/v1"
	"GoBlog/middleware"
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{

		auth.PUT("user/:id", v1.EditUser)

		auth.DELETE("user/:id", v1.DeleteUser)

		// 分类的模块的路由接口
		auth.POST("category/add", v1.AddCate)

		auth.PUT("category/:id", v1.EditCate)

		auth.DELETE("category/:id", v1.DeleteCate)

		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)

		auth.PUT("article/:id", v1.EditArticle)

		auth.DELETE("article/:id", v1.DeleteArticle)

		// 上传文件
		auth.POST("upload", v1.Upload)
	}

	public := r.Group("api/v1")
	{
		// User模块的路由接口
		public.POST("user/add", v1.AddUser)

		public.GET("users", v1.GetUsers)

		public.GET("category", v1.GetCate)

		public.GET("article", v1.GetArt)

		public.GET("article/list/:id", v1.GetCatArt)

		public.GET("article/info/:id", v1.GetArtInfo)

		public.POST("login", v1.Login)

	}

	r.Run(utils.HttpPort)
}
