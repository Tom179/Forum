package mail

type Driver interface { //driver_interface是为了后续扩展使用其他发送邮件的渠道提供了方便。
	// 检查验证码
	Send(email Email, config map[string]string) bool
}
