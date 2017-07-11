package aliMail

import (
	"net/smtp"
	"strings"
)

//一次性发送多个邮件或者单个邮件
func (g *Gun) SendMail(c Cont) (err error) {
	to := strings.Split(c.Aims, ";")
	msg := []byte("To: " + c.Aims + "\r\nFrom: " + g.User + "\r\nSubject: " + c.Subject + "\r\n" + c.Content_type + "\r\n\r\n" + c.Body)
	err = smtp.SendMail(g.Host, g.Auth, g.User, to, msg)
	return
}
