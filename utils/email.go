package utils

import (
	"ginserver/global"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	mailConn := map[string]string{
		"user": global.AppSetting.Email.Sender,
		"pass": global.AppSetting.Email.Password,
		"host": global.AppSetting.Email.Host,
		"port": global.AppSetting.Email.Port,
	}

	port, _ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "UTOKIA"))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}
