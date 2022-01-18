package aggregate

import (
	"encoding/json"
	"fmt"
	"sap_cert_mgt/domain/enums"
	"sap_cert_mgt/domain/repository"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/logs"
	"sap_cert_mgt/infra/model"
	"sap_cert_mgt/infra/persistence"
	"net/smtp"
	"strings"
	"sync"
	"time"
)

//go:generate mockgen -source=./email_notice.go -destination ./mock/mock_email_notice.go -package mock
type EmailNoticeEntityInterface interface {
	Scan() (err error)
	Send() (err error)
}

var (
	emailNoticeServiceOnce sync.Once
	emailNoticeService     EmailNoticeEntityInterface
)

type emailNoticeEntity struct {
	certRepo           repository.CertRepoInterface           // 依赖抽象
	noticeStrategyRepo repository.NoticeStrategyRepoInterface // 依赖抽象
	noticeConfRepo     repository.NoticeConfRepoInterface     // 依赖抽象
	noticeEventRepo    repository.NoticeEventRepoInterface    // 依赖抽象
	log                logs.Logger
}

var _ EmailNoticeEntityInterface = &emailNoticeEntity{}

func NewEmailNoticeEntity() EmailNoticeEntityInterface {
	emailNoticeServiceOnce.Do(func() {
		emailNoticeService = &emailNoticeEntity{
			certRepo:           persistence.NewCertRepo(),
			noticeStrategyRepo: persistence.NewNoticeStrategyRepo(),
			noticeConfRepo:     persistence.NewNoticeConfRepo(),
			noticeEventRepo:    persistence.NewNoticeEventRepo(),
			log:                logs.NewLogger(),
		}
	})
	return emailNoticeService
}

func (service *emailNoticeEntity) Scan() (err error) {
	// 获取所有启用的策略(NoticeType TriggerThreshold ToEmails)
	filterStrategy := make(map[string]interface{})
	filterStrategy["limit"] = 100
	filterStrategy["enabled_state"] = enums.Enabled
	filterStrategy["deleted_state"] = enums.NotDeleted
	_, strategyList, err := service.noticeStrategyRepo.GetList(filterStrategy)
	if err != nil {
		return
	}
	if len(strategyList) == 0 {
		return
	}
	now := time.Now()
	for _, strategy := range strategyList {
		// 获取启用状态下所有临期的证书(NotAfter EnabledState)
		td, _ := time.ParseDuration(fmt.Sprintf("%dh", 24*strategy.TriggerThreshold))
		noticeTime := now.Add(td)
		filterCert := make(map[string]interface{})
		filterCert["limit"] = 1
		filterCert["enabled_state"] = enums.Enabled
		filterCertArgs := make([]interface{}, 0)
		filterCertArgs = append(filterCertArgs, "not_after <= ?")
		filterCertArgs = append(filterCertArgs, noticeTime)
		_, certList, _ := service.certRepo.GetList(filterCert, filterCertArgs...)
		if len(certList) == 0 {
			continue
		}
		certInfo := certList[0]
		//notAfter := certInfo.NotAfter
		// 跳过已经处理过的通知事件（去重，单个策略只成功推送一次）
		filterEvent := make(map[string]interface{})
		filterEvent["limit"] = 1000
		filterEvent["cert_id"] = certInfo.Id
		filterEvent["notice_strategy_id"] = strategy.Id
		filterEventArgs := make([]interface{}, 0)
		filterEventArgs = append(filterEventArgs, "event_state in ?")
		filterEventArgs = append(filterEventArgs, []int{enums.EventStateWaiting, enums.EventStateProcess, enums.EventStateSuccess})
		_, eventList, _ := service.noticeEventRepo.GetList(filterEvent, filterEventArgs...)
		if len(eventList) > 0 {
			continue
		}
		// 保存通知事件，并设置状态为 waiting
		eventInfo := &model.NoticeEvent{}
		eventInfo.CertId = certInfo.Id
		eventInfo.NoticeStrategyId = strategy.Id
		eventInfo.EventState = enums.EventStateWaiting
		_, err = service.noticeEventRepo.Create(eventInfo)
	}

	return
}

func (service *emailNoticeEntity) Send() (err error) {
	// 获取邮件smtp配置，为空报错退出
	noticeConf, err := service.noticeConfRepo.GetEmail()
	if err != nil {
		return
	}
	// 响应处理
	noticeConfEmail := &vo.NoticeConfGetEmailRes{}
	if noticeConf.ConfigData == "" {
		return
	}
	err = json.Unmarshal([]byte(noticeConf.ConfigData), &noticeConfEmail)
	if err != nil {
		return
	}
	// 获取所有通知事件(CertId NoticeStrategyId EventState)
	filterEvent := make(map[string]interface{})
	filterEvent["limit"] = 1000
	filterEvent["event_state"] = enums.EventStateWaiting
	_, eventList, err := service.noticeEventRepo.GetList(filterEvent)
	if err != nil {
		return
	}
	for _, eventInfo := range eventList {
		// 获取关联证书信息(NotAfter)
		certInfo, err := service.certRepo.GetInfo(eventInfo.CertId)
		if err != nil {
			continue
		}
		// 获取关联策略信息(NoticeType TriggerThreshold ToEmails)
		strategyInfo, err := service.noticeStrategyRepo.GetInfo(eventInfo.NoticeStrategyId)
		if err != nil {
			continue
		}
		// 拼装邮件信息，发送通知邮件
		err = service.SendEmail(
			noticeConfEmail,
			strategyInfo.ToEmails,
			"证书过期提醒",
			fmt.Sprintf("证书：%s，即将在：%s过期，请联系SAP管理员更新证书", certInfo.AuthId, certInfo.NotAfter),
		)
		eventData := make(map[string]interface{})
		if err != nil {
			service.log.Errorln("send email failure")
			eventData["event_tate"] = enums.EventStateFailure
			_ = service.noticeEventRepo.Update(eventInfo.Id, eventData)
		} else {
			service.log.Infoln("send email success")
			eventData["event_tate"] = enums.EventStateSuccess
			_ = service.noticeEventRepo.Update(eventInfo.Id, eventData)
		}
	}
	return
}

func (service *emailNoticeEntity) SendEmail(ec *vo.NoticeConfGetEmailRes, receivers, subject, body string) (err error) {
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
	err = smtp.SendMail(serverAddress, auth, ec.FromEmail, sendTo, msg)
	return
}
