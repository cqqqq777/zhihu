package boot

import (
	"github.com/gin-gonic/gin"
	g "zhihu/global"
	"zhihu/middleware"
)

func InitRouters() {
	r := gin.New()
	r.Use(middleware.GinLogger(g.Logger), middleware.GinRecovery(g.Logger, true))
}
