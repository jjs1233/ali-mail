package aliMail

type Cont struct {
	Aims         string
	Subject      string
	Body         string
	Mailtype     string
	Content_type string
}

func NewCont(aims string, subject string, mailtype string, body string) Cont {

	var content_type string

	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	return Cont{
		Aims:         aims,
		Subject:      subject,
		Body:         body,
		Mailtype:     mailtype,
		Content_type: content_type,
	}
}
