package Email

import (
	"gopkg.in/gomail.v2"
	"gvb_server/core"
	"gvb_server/global"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "警告通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}
func NewCode() Api {
	return Api{Subject: Code}
}
func NewNote() Api {
	return Api{Subject: Note}
}
func NewAlarm() Api {
	return Api{Subject: Alarm}
}
func send(name, subject, body string) error {
	e := global.Config.Email
	return SendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}

func SendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName)) //发件人邮箱，发件人名字
	m.SetHeader("To", mailTo)                                //发送给谁
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}

func main() {
	core.InitConf()
	core.InitLogger()
	NewCode().Send("2371522043@qq.com", "验证码是2056")
}
