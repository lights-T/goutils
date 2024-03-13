package email

import (
	"fmt"
	"testing"

	lfiles "github.com/lights-T/goutils/files"
)

func Test_Running(t *testing.T) {
	content := fmt.Sprintf(
		"****************************************************************************************************</br>" +
			">>>>>>>>>>>>>>>>>>>>>>> 如果您对本邮件有任何意见/要求，请联系XXX@whchem.com <<<<<<<<<<<<<<<<<<<<<<<</br>" +
			"****************************************************************************************************</br>" +
			"您好,</br></br>" +
			"您有新的培训计划。</br>" +
			"这是一封自动发送的邮件，请不要回复此邮件。</br></br>" +
			"谢谢！",
	)
	conf := &Conf{
		UserName: "xxx@qq.com",
		Password: "xxx",
		Host:     "smtp.qq.com",
		Port:     587,
		From:     "xxx@qq.com",
		To:       []string{"yyy@yyy.com"},
		Subject:  "邮件发送测试",
		Body:     content,
		//Body: "********************************************************************************************************\n\n\n\n\n\n\n\n\n\n" +
		//	">>>>>>If you have any comments/requests regarding to this email, please contact xxx@xxx.com <<<<<<\n" +
		//	"********************************************************************************************************\n" +
		//	"Dear Sir, \n" +
		//	"These are the alarm daily reports for today. \nThis is an auto-send email, please don't reply this email. \n" +
		//	"Thank you!",
	}
	e := NewEmail(conf)
	file := "./att.txt"
	exist := lfiles.FileExist(file)
	if !exist {
		t.Fatalf("File %s does not exist", file)
	}

	e.EmailMsg.Attach(file)
	if err := e.EmailConn.DialAndSend(e.EmailMsg); err != nil {
		t.Fatalf("Mail delivery failure, err: %s", err.Error())
	}
	t.Log("Done.")
}
