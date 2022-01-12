package clients

import (
	"net/smtp"
	"strings"
)

type EmailNoticeConfig struct {
	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerHost string `yaml:"server_host"`
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465，其它默认25
	ServerPort string `yaml:"server_port"`
	// FromName　发件人邮箱名称，可以为空
	FromName string `yaml:"from_name"`
	// FromEmail　发件人邮箱地址
	FromEmail string `yaml:"from_email"`
	// FromPasswd 发件人邮箱客户端授权码（需要开通）注意：不是密码
	FromPasswd string `yaml:"from_passwd"`
}

type EmailNoticeClient struct {
	EmailNoticeConfig
}

// 实现依赖倒置（检查接口是否实现）
var _ NoticeInterface = &EmailNoticeClient{}

// Send 邮件发送
func (ec *EmailNoticeClient) Send(receivers, subject, body string) error {
	auth := smtp.PlainAuth("", ec.FromEmail, ec.FromPasswd, ec.ServerHost)
	contentType := "Content-Type: text/html; charset=UTF-8"
	sendFrom := ec.FromEmail
	if ec.FromName != "" {
		sendFrom = ec.FromName + "<" + ec.FromEmail + ">"
	}
	var data []string
	data = append(data, "From: "+sendFrom)
	data = append(data, "To: "+receivers)
	data = append(data, "Subject: "+subject)
	data = append(data, contentType)
	data = append(data, "")
	data = append(data, body)
	msg := []byte(strings.Join(data, "\r\n"))
	sendTo := strings.Split(receivers, ",")
	serverAddress := strings.Join([]string{ec.ServerHost, ec.ServerPort}, ":")
	err := smtp.SendMail(serverAddress, auth, ec.FromEmail, sendTo, msg)
	return err
}

func NewEmailNoticeClient(emailNoticeConfig EmailNoticeConfig) (*EmailNoticeClient, error) {
	emailClient := &EmailNoticeClient{
		emailNoticeConfig,
	}
	return emailClient, nil
}
