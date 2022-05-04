package entity

import (
	"sap_cert_mgt/domain/repository"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/model"
	"sap_cert_mgt/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./notice_strategy.go -destination ./mock/mock_notice_strategy.go -package mock
type NoticeStrategyEntityInterface interface {
	AddNoticeStrategy(data *vo.NoticeStrategyCreateReq, createdBy string) (id int, err error)
	DelNoticeStrategy(id int, deletedBy string) (err error)
	ModNoticeStrategy(id int, data map[string]interface{}, updatedBy string) (err error)
	GetNoticeStrategyInfo(id int) (data *vo.NoticeStrategyGetInfoRes, err error)
	GetNoticeStrategyList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.NoticeStrategyGetInfoRes, err error)
}

var (
	noticeStrategyServiceOnce sync.Once
	noticeStrategyService     NoticeStrategyEntityInterface
)

type noticeStrategyEntity struct {
	noticeStrategyRepo repository.NoticeStrategyRepoInterface // 依赖抽象
	opLogRepo          repository.OperationLogRepoInterface   // 依赖抽象
}

var _ NoticeStrategyEntityInterface = &noticeStrategyEntity{}

func NewNoticeStrategyEntity() NoticeStrategyEntityInterface {
	noticeStrategyServiceOnce.Do(func() {
		noticeStrategyService = &noticeStrategyEntity{
			noticeStrategyRepo: persistence.NewNoticeStrategyRepo(),
			opLogRepo:          persistence.NewOperationLogRepo(),
		}
	})
	return noticeStrategyService
}

func (service *noticeStrategyEntity) AddNoticeStrategy(data *vo.NoticeStrategyCreateReq, createdBy string) (id int, err error) {
	// 参数处理
	confInfo := &model.NoticeStrategy{}
	confInfo.NoticeType = *data.NoticeType
	confInfo.TriggerThreshold = *data.TriggerThreshold
	confInfo.ToEmails = data.ToEmails
	confInfo.EnabledState = *data.EnabledState

	id, err = service.noticeStrategyRepo.Create(confInfo, createdBy)
	if err != nil {
		return
	}
	// 操作日志
	opLogData := &model.OperationLog{
		OpType:   "create",
		RsType:   "notice_strategy",
		RsId:     id,
		OpDetail: "",
		OpError:  "",
	}
	_, err = service.opLogRepo.Create(opLogData, createdBy)
	return
}

func (service *noticeStrategyEntity) ModNoticeStrategy(id int, data map[string]interface{}, updatedBy string) (err error) {
	err = service.noticeStrategyRepo.Update(id, data, updatedBy)
	if err != nil {
		return
	}
	// 操作日志
	opLogData := &model.OperationLog{
		OpType:   "update",
		RsType:   "notice_strategy",
		RsId:     id,
		OpDetail: "",
		OpError:  "",
	}
	_, err = service.opLogRepo.Create(opLogData, updatedBy)
	return
}

func (service *noticeStrategyEntity) DelNoticeStrategy(id int, deletedBy string) (err error) {
	err = service.noticeStrategyRepo.Delete(id, deletedBy)
	if err != nil {
		return
	}
	// 操作日志
	opLogData := &model.OperationLog{
		OpType:   "delete",
		RsType:   "notice_strategy",
		RsId:     id,
		OpDetail: "",
		OpError:  "",
	}
	_, err = service.opLogRepo.Create(opLogData, deletedBy)
	return
}

func (service *noticeStrategyEntity) GetNoticeStrategyInfo(id int) (data *vo.NoticeStrategyGetInfoRes, err error) {
	resInfo, err := service.noticeStrategyRepo.GetInfo(id)
	// 响应处理
	data = &vo.NoticeStrategyGetInfoRes{}
	data.Id = resInfo.Id
	data.NoticeType = resInfo.NoticeType
	data.TriggerThreshold = resInfo.TriggerThreshold
	data.EnabledState = resInfo.EnabledState
	data.ToEmails = resInfo.ToEmails
	data.SetNoticeTypeDisplayName()
	data.SetEnabledStateDisplayName()
	return
}

func (service *noticeStrategyEntity) GetNoticeStrategyList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.NoticeStrategyGetInfoRes, err error) {
	total, resList, err := service.noticeStrategyRepo.GetList(filter, args...)
	// 响应处理
	data = make([]*vo.NoticeStrategyGetInfoRes, 0)
	for _, resInfo := range resList {
		item := &vo.NoticeStrategyGetInfoRes{}
		item.Id = resInfo.Id
		item.NoticeType = resInfo.NoticeType
		item.TriggerThreshold = resInfo.TriggerThreshold
		item.EnabledState = resInfo.EnabledState
		item.ToEmails = resInfo.ToEmails
		item.CreatedAt = resInfo.CreatedAt
		item.CreatedBy = resInfo.CreatedBy
		item.UpdatedAt = resInfo.UpdatedAt
		item.UpdatedBy = resInfo.UpdatedBy
		item.SetNoticeTypeDisplayName()
		item.SetEnabledStateDisplayName()
		data = append(data, item)
	}
	return
}
