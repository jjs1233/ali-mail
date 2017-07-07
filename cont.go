package aliMail

type Cont struct {
	Aims         string
	Subject      string
	Body         string
	Mailtype     string
	Content_type string
}

func (c *Cont) init() {
	if c.Mailtype == "html" {
		c.Content_type = "Content-Type: text/" + c.Mailtype + "; charset=UTF-8"
	} else {
		c.Content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
}
