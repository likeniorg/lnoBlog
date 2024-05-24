package main

import (
	"net/smtp"
)

func sendMail(toEmail string, msg []byte) error {
	smtpHost := "localhost" // 替换为你的SMTP服务器地址
	smtpPort := "25"        // 替换为你的SMTP服务器端口

	from := "kk@kkkjjj.likeni.org"
	tos := []string{toEmail}

	// 使用 nil 作为 auth 参数表示不需要身份验证
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, tos, msg)
	if err != nil {
		return err
	}
	return nil
}
