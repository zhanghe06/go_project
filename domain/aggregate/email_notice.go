package aggregate

import (
	"go_project/domain/repository"
	"go_project/infra/persistence"
	"sync"
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
	certRepo repository.CertRepoInterface // 依赖抽象
	noticeStrategyRepo repository.NoticeStrategyRepoInterface // 依赖抽象
	noticeConfRepo repository.NoticeConfRepoInterface // 依赖抽象
	noticeEventRepo repository.NoticeEventRepoInterface // 依赖抽象
}

var _ EmailNoticeEntityInterface = &emailNoticeEntity{}

func NewEmailNoticeEntity() EmailNoticeEntityInterface {
	emailNoticeServiceOnce.Do(func() {
		emailNoticeService = &emailNoticeEntity{
			certRepo: persistence.NewCertRepo(),
			noticeStrategyRepo: persistence.NewNoticeStrategyRepo(),
			noticeConfRepo: persistence.NewNoticeConfRepo(),
			noticeEventRepo: persistence.NewNoticeEventRepo(),
		}
	})
	return emailNoticeService
}

func (service *emailNoticeEntity) Scan() (err error) {
	// todo
	// 获取所有启用的策略(NoticeType TriggerThreshold ToEmails)
	// 获取启用状态下所有临期的证书(NotAfter EnabledState)
	// 保存通知事件，并设置状态为 waiting
	return
}

func (service *emailNoticeEntity) Send() (err error) {
	// todo
	// 获取邮件smtp配置，为空报错退出
	// 获取所有通知事件(CertId NoticeStrategyId EventState)
	// 获取关联证书信息(NotAfter)
	// 获取关联策略信息(NoticeType TriggerThreshold ToEmails)
	// 拼装邮件信息，发送通知邮件
	return
}
