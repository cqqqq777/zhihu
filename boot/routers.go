package boot

import (
	"github.com/gin-gonic/gin"
	"zhihu/controller"
	g "zhihu/global"
	"zhihu/middleware"
)

func InitRouters() {
	r := gin.New()
	r.Use(middleware.GinLogger(g.Logger), middleware.GinRecovery(g.Logger, true))
	r.Use(middleware.Cors)
	v1 := r.Group("/api/v1")
	public := v1.Group("")
	{
		public.POST("/registration", controller.Register)
		public.POST("/verification", controller.PostVerification)
		public.POST("/login", controller.Login)
		public.PUT("/password/forget", controller.ForgetPassword)
		public.GET("/topics", controller.GetAllTopic)
	}
	private := v1.Group("")
	private.Use(middleware.JWTAuth)
	{
		private.PUT("/password", controller.RevisePassword)
		private.PUT("/username", controller.ReviseUsername)
	}
	if err := r.Run(); err != nil {
		panic(err)
	}
}
