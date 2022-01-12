package entity

import (
	"encoding/json"
	"go_project/domain/repository"
	"go_project/domain/vo"
	"go_project/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./notice_conf.go -destination ./mock/mock_notice_conf.go -package mock
type NoticeConfEntityInterface interface {
	GetNoticeConfEmail() (data *vo.NoticeConfGetEmailRes, err error)
	ModNoticeConfEmail(data map[string]interface{}, updatedBy string) (err error)
}

var (
	noticeConfServiceOnce sync.Once
	noticeConfService     NoticeConfEntityInterface
)

type noticeConfEntity struct {
	noticeConfRepo repository.NoticeConfRepoInterface // 依赖抽象
}

var _ NoticeConfEntityInterface = &noticeConfEntity{}

func NewNoticeConfEntity() NoticeConfEntityInterface {
	noticeConfServiceOnce.Do(func() {
		noticeConfService = &noticeConfEntity{
			noticeConfRepo: persistence.NewNoticeConfRepo(),
		}
	})
	return noticeConfService
}

func (service *noticeConfEntity) GetNoticeConfEmail() (data *vo.NoticeConfGetEmailRes, err error) {
	noticeConf, _ := service.noticeConfRepo.GetEmail()
	// 响应处理
	data = &vo.NoticeConfGetEmailRes{}
	if noticeConf.ConfigData == "" {
		return
	}
	err = json.Unmarshal([]byte(noticeConf.ConfigData), &data)
	return
}

func (service *noticeConfEntity) ModNoticeConfEmail(data map[string]interface{}, updatedBy string) (err error) {
	return service.noticeConfRepo.ModEmail(data, updatedBy)
}
