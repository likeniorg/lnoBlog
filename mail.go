package main

import (
	"net/smtp"
)

func sendMail(toEmail string, subject string, msg string) error {
	smtpHost := "localhost" // 替换为你的SMTP服务器地址
	smtpPort := "25"        // 替换为你的SMTP服务器端口

	from := "kk@kkkjjj.likeni.org"
	tos := []string{toEmail}
	content := []byte("Subject: " + subject + "\r\n\r\n" + msg)
	// 使用 nil 作为 auth 参数表示不需要身份验证
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, tos, content)
	if err != nil {
		return err
	}
	return nil
}
