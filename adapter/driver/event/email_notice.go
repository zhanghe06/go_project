package event

import (
	"go_project/domain/aggregate"
	"go_project/infra/logs"
	"sync"
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

type emailNotice struct {
	//EmailNoticeConfig
	emailNoticeEntity aggregate.EmailNoticeEntityInterface
	log logs.Logger
}

var (
	noticeOnce sync.Once
	notice NoticeInterface
)

func NewEmailNotice() NoticeInterface {
	noticeOnce.Do(func() {
		notice = &emailNotice{
			emailNoticeEntity: aggregate.NewEmailNoticeEntity(),
			log:              logs.NewLogger(),
		}
	})
	return notice
}


// 实现依赖倒置（检查接口是否实现）
var _ NoticeInterface = &emailNotice{}

func (ec *emailNotice) Scan() {
	err := ec.emailNoticeEntity.Scan()
	if err != nil {
		ec.log.Errorln(err)
	}
}

func (ec *emailNotice) Send() {
	err := ec.emailNoticeEntity.Send()
	if err != nil {
		ec.log.Errorln(err)
	}
}

// Send 邮件发送
//func (ec *emailNotice) Send(receivers, subject, body string) error {
//	auth := smtp.PlainAuth("", ec.FromEmail, ec.FromPasswd, ec.ServerHost)
//	contentType := "Content-Type: text/html; charset=UTF-8"
//	sendFrom := ec.FromEmail
//	if ec.FromName != "" {
//		sendFrom = ec.FromName + "<" + ec.FromEmail + ">"
//	}
//	var data []string
//	data = append(data, "From: "+sendFrom)
//	data = append(data, "To: "+receivers)
//	data = append(data, "Subject: "+subject)
//	data = append(data, contentType)
//	data = append(data, "")
//	data = append(data, body)
//	msg := []byte(strings.Join(data, "\r\n"))
//	sendTo := strings.Split(receivers, ",")
//	serverAddress := strings.Join([]string{ec.ServerHost, ec.ServerPort}, ":")
//	err := smtp.SendMail(serverAddress, auth, ec.FromEmail, sendTo, msg)
//	return err
//}

