package entity

import (
	"encoding/json"
	"sap_cert_mgt/domain/repository"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/model"
	"sap_cert_mgt/infra/persistence"
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
	noticeConfRepo repository.NoticeConfRepoInterface   // 依赖抽象
	opLogRepo      repository.OperationLogRepoInterface // 依赖抽象
}

var _ NoticeConfEntityInterface = &noticeConfEntity{}

func NewNoticeConfEntity() NoticeConfEntityInterface {
	noticeConfServiceOnce.Do(func() {
		noticeConfService = &noticeConfEntity{
			noticeConfRepo: persistence.NewNoticeConfRepo(),
			opLogRepo:      persistence.NewOperationLogRepo(),
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
	err = service.noticeConfRepo.ModEmail(data, updatedBy)
	if err != nil {
		return
	}
	// 操作日志
	opLogData := &model.OperationLog{
		OpType:   "update",
		RsType:   "notice_conf",
		RsId:     0,
		OpDetail: "update email notice conf ",
		OpError:  "",
	}
	_, err = service.opLogRepo.Create(opLogData, updatedBy)
	return
}
