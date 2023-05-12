package mail

import (
	"fmt"
	"net/smtp"

	emailPKG "github.com/jordan-wright/email"
)

// SMTP 实现 email.Driver interface
type SMTP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SMTP) Send(email Email, config map[string]string) bool { //传入Email结构体

	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	//logger.DebugJSON("发送邮件", "发件详情", e)//日志
	fmt.Println("发送邮件：", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),

		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)
	if err != nil {
		//logger.ErrorString("发送邮件", "发件出错", err.Error())//日志
		fmt.Println("发送邮件错误", err.Error())
		return false
	}

	//logger.DebugString("发送邮件", "发件成功", "")//日志
	fmt.Println("发送邮件成功")
	return true
}
