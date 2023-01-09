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
		public.GET("/users/:uid/info", controller.GetUserInfo)
		public.GET("/topics", controller.GetAllTopic)
		public.GET("/topics/:tid", controller.TopicDetail)
		public.GET("/posts/:pid", controller.PostDetail)
		public.GET("/questions", controller.QuestionList)
		public.GET("/essays", controller.EssayList)
		public.GET("/posts", controller.SearchPost)

	}
	private := v1.Group("")
	private.Use(middleware.JWTAuth)
	{
		private.PUT("/password", controller.RevisePassword)
		private.PUT("/username", controller.ReviseUsername)
		private.PUT("/users/:uid/info", controller.UpdateUserInfo)
		private.POST("/post", controller.CreatePost)
		private.GET("/user/questions", controller.UserQuestionList)
		private.GET("/user/essays", controller.UserEssayList)
		private.PUT("/posts/:pid", controller.UpdatePost)
		private.DELETE("/posts/:pid", controller.DeletePost)
		private.POST("/posts/:pid/star", controller.StarPost)
	}
	if err := r.Run(); err != nil {
		panic(err)
	}
}
