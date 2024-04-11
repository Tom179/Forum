package requests //请求验证容错

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goWeb/app/requests/validators"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"` //valid是一个常用的第三方验证库，可以定义结构体字段的验证规则，如果不符合，验证库会返回相应的错误信息
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	//request包下的SignupPhoneExist仅定义格式：后续可换函数名
	//传入请求结构体，返回ValidateStruct函数
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"}, //不为空的11位字符串
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	// 配置初始化
	return validate(data, rules, messages)
}

// SignupUsingEmailRequest 通过邮箱注册的请求信息
type SignupUsingEmailRequest struct {
	Email           string `json:"email,omitempty" valid:"email"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}

func SignupUsingEmail(data interface{}, c *gin.Context) map[string][]string {
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email"},  //, "not_exists:users,email"不能写
		"name":             []string{"required", "alpha_num", "between:3,20"}, //, "not_exists:users,name"不能写
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingEmailRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}
