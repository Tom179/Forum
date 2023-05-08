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
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}

	}
}
