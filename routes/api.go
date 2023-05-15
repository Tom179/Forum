package routes

import (
	"github.com/gin-gonic/gin"
	"goWeb/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegistAPIRouters(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "goWEb!!!",
			})
		})
		authGroup := v1.Group("/auth") //认证组
		{
			suc := new(auth.SignupController) //创建对象，该对象封装了函数
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			//👆查询手机号是否存在接口，传入的是查询手机号是否存在的处理函数（传入c *gin.Context），该函数再封装的查询工具函数
			authGroup.POST("signup/email/exist", suc.IsEmailExist)

			vcc := new(auth.VerifyCodeController)
			//创建VerifyCodeController的实例
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha) //图片验证码接口
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)

			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
		}
	}
}
