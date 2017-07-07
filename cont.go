package aliMail

type Cont struct {
	Aims         string `收件人 多个采用 ; 分开`
	Subject      string `邮件标题`
	Body         string `邮件正文`
	Mailtype     string `邮件类型 html`
	Content_type string
}

func (c *Cont) init() {
	if c.Mailtype == "html" {
		c.Content_type = "Content-Type: text/" + c.Mailtype + "; charset=UTF-8"
	} else {
		c.Content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
}
