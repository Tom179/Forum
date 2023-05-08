package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "goWeb/app/http/controllers/api/v1"
	"goWeb/app/models/user"
	"net/http"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) { //处理函数
	type PhoneExistRequest struct {
		Phone string `json:"phone"` //与json格式转换
	}
	request := PhoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil { //从请求c中绑定参数到request结构体中
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{ //如果成功获取到请求，验证
		"exist": user.IsPhoneExist(request.Phone),
	})

}
