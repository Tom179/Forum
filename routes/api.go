package routes

import (
	"github.com/gin-gonic/gin"
	"goWeb/app/http/MiddleWares"
	controllers "goWeb/app/http/controllers/api/v1"
	"goWeb/app/http/controllers/api/v1/auth"
)

func RegistAPIRouters(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Air重载测试",
		})
	})
	v1.Use(MiddleWares.LimitIP("200-H")) //全局中间件:每小时200个请求？
	{
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

			//创建LoginController的实例
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-password", lgc.LoginByPassword)

			authGroup.POST("/login/refresh-token", lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			//authGroup.POST("/password-reset/using-phone", pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", pwc.ResetByEmail)

			uc := new(controllers.UsersController)
			// 获取当前用户
			v1.GET("/user", MiddleWares.AuthJWT(), uc.CurrentUser)
			usersGroup := v1.Group("/users")
			{
				usersGroup.GET("", uc.Index)
			}

			v1.GET("/testIp", func(context *gin.Context) {
				ip := context.ClientIP()
				context.JSON(200, gin.H{"": context.FullPath() + ip})
			})
			cgc := new(controllers.CategoriesController)
			cgcGroup := v1.Group("/categories")
			{
				cgcGroup.GET("", cgc.Index)

				cgcGroup.POST("", MiddleWares.AuthJWT(), cgc.Store)
				cgcGroup.PUT("/:id", MiddleWares.AuthJWT(), cgc.Update)
				cgcGroup.DELETE("/:id", MiddleWares.AuthJWT(), cgc.Delete)
			}

		}

	}
}
