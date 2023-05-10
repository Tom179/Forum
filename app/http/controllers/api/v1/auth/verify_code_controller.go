package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "goWeb/app/http/controllers/api/v1"
	"goWeb/app/requests"
	"goWeb/app/response"
	"goWeb/captcha"
	"goWeb/verifycode"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha() //1.先是绑定到redis库中2.生成id,验证码的值并存储到redis中，最后生成图片，返回string中的后段id、和图片的base64编码
	//logger.LogIf(err)// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	if err != nil {
		fmt.Println(err)
	}
	// 返回给用户
	/*c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})*/
	response.JSON(c, gin.H{ //简单封装
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// SendUsingEmail 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	// 2. 发送邮件
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送 Email 验证码失败~")
	} else {
		response.Success(c)
	}
}
