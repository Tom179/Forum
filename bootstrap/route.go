package bootstrap

import (
	"github.com/gin-gonic/gin"
	"goWeb/routes"
)

func SetupRoute(r *gin.Engine) { //传入engine
	registGlobalMiddleWare(r)  //注册中间件
	routes.RegistAPIRouters(r) //开启路由

}
func registGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}
