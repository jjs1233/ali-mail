package main

import (
	"fmt"

	"github.com/ali-mail/aliMail"
)

func main() {
	/*
		阿里云邮件推送服务测试
	*/

	g := aliMail.NewGun(&aliMail.GunConfig{
		User:     "coobii@mail-test.coobii.com",
		Password: "QWerty1234",
		Host:     "smtpdm.aliyun.com:25",
	})

	html := `
        <html>
        <body>
        <h3>
        "您的验证码为 XXXXX"
        </h3>
        </body>
        </html>
        `
	err := g.SendMail(&aliMail.Cont{
		Aims:     "jjs1233@163.com;www.1402960567@qq.com",
		Subject:  "这是次测试",
		Body:     html,
		Mailtype: "html",
	})

	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}

}
