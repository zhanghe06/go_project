package clients

import (
	"testing"
)

type Config struct{
	EmailNotice EmailNoticeConfig `yaml:"email_notice"`
}

func TestEmail(t *testing.T) {
	// 读取配置文件
	//configFile, err := ioutil.ReadFile("email_test.yaml")
	//if err != nil {
	//	log.Panicf(err.Error())
	//}
	//var config Config
	//if err = yaml.Unmarshal(configFile, &config); err != nil {
	//	log.Panicf(err.Error())
	//}
	//client, err := NewEmailNoticeClient(config.EmailNotice)
	//if err != nil {
	//	log.Panicf(err.Error())
	//}
	//receivers := "xxxxxx@163.com"
	//subject := "使用Golang发送邮件"
	//
	//body := `
	//	<html>
	//	<body>
	//	<h3>
	//	"Test send to email"
	//	</h3>
	//	</body>
	//	</html>
	//	`
	//err = client.Send(receivers, subject, body)
	//if err != nil {
	//	log.Panicf(err.Error())
	//}
}