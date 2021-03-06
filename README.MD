# golang 实现阿里云smtp邮件推送接口调用

[![Build Status](https://api.travis-ci.org/jjs1233/aliMail.svg?branch=master)](https://travis-ci.org/jjs1233/aliMail)

    go get "github.com/jjs1233/aliMail"

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

    c := aliMail.NewCont(
      Aims      =  "收件人 多个采用 ; 分开",
      Subject   =  "标题",
      Mailtype  =  "html",
      Body      =  html,
    )
    
    err := g.SendMail(c)

    if err != nil {
      fmt.Println("Send mail error!")
      fmt.Println(err)
    } else {
      fmt.Println("Send mail success!")
    }
}
```

### 数据类型

#### GunConfig

```go
type GunConfig struct {
	User     string  "控制台创建的发信地址",
	Password string  "控制台设置的SMTP密码",
	Host     string  "smtpdm.aliyun.com:25",
}

```


#### Gun

```go
type GunConfig struct {
	GunConfig 
	Auth smtp.Auth
}

```

### Cont
```go
type Cont struct {
	Aims         string  "收件人 多个采用 ; 分开"
	Subject      string  "邮件标题"
	Body         string  "邮件正文"
	Mailtype     string  "邮件类型 html"
	Content_type string
}
```
