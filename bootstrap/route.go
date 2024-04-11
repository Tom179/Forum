package bootstrap

import (
	"github.com/gin-gonic/gin"
	"goWeb/app/http/MiddleWares"
	"goWeb/routes"
)

func SetupRoute(r *gin.Engine) { //传入engine
	registGlobalMiddleWare(r)  //注册中间件
	routes.RegistAPIRouters(r) //开启路由
}
func registGlobalMiddleWare(r *gin.Engine) {
	r.Use( //注册全局中间件到r引擎上，本来不这样写就会默认添加logger()和recovery()函数，这样写是为了方便后续添加其他中间件
		//gin.Logger(), //gin内置的日志中间件，我们更换为自定义的zap()
		MiddleWares.Logger(),
		gin.Recovery(),
	)

}
