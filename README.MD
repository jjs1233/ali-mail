# golang 实现阿里大于smtp邮件推送接口调用

[![Build Status](https://api.travis-ci.org/jjs1233/aliMail.svg?branch=master)](https://travis-ci.org/jjs1233/aliMail)

    import "github.com/jjs1233/aliMail"

## Usage

### 用法

```go
package main

import (
    "github.com/jjs1233/aliMail"
)

func main() {
    g := aliMail.NewGun(&aliMail.GunConfig{
      User:     "控制台创建的发信地址",
      Password: "控制台设置的SMTP密码",
      Host:     "smtpdm.aliyun.com:25",
    })
    fmt.Println(g.Auth)
    /*发件内容*/
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
      Aims:     "收件人 多个采用 ; 分开",
      Subject:  "标题",
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
