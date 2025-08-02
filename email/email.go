package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Service struct {
	EmailConn *gomail.Dialer
	EmailMsg  *gomail.Message
	EmailConf *Conf
}

type Conf struct {
	UserName       string `json:"userName"`
	Password       string `json:"password"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	From           string `json:"from"`
	RecordFilePath string
	To             []string
	Bcc            []string
	Cc             []string
	Subject        string
	Body           string // Html message (optional)
	SecureVerify   bool   //是否验证服务器证书等信息，默认不验证
}

func NewEmail(c *Conf) *Service {
	m := gomail.NewMessage()
	m.SetHeader("From", c.From)
	m.SetHeader("To", c.To...)
	for _, v := range c.Cc {
		m.SetAddressHeader("Cc", v, v)
	}
	for _, v := range c.Bcc {
		m.SetAddressHeader("Bcc", v, v)
	}
	m.SetHeader("Subject", c.Subject)
	m.SetBody("text/html", c.Body)
	emailConn := gomail.NewDialer(c.Host, c.Port, c.UserName, c.Password)
	emailConn.TLSConfig = &tls.Config{InsecureSkipVerify: !c.SecureVerify}
	return &Service{
		EmailConn: emailConn,
		EmailConf: c,
		EmailMsg:  m,
	}
}
