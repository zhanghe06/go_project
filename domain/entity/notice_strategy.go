package entity

import (
	"go_project/domain/repository"
	"go_project/domain/vo"
	"go_project/infra/model"
	"go_project/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./notice_strategy.go -destination ./mock/mock_notice_strategy.go -package mock
type NoticeStrategyEntityInterface interface {
	AddNoticeStrategy(data *vo.NoticeStrategyCreateReq) (id int, err error)
	DelNoticeStrategy(id int) (err error)
	ModNoticeStrategy(id int, data map[string]interface{}) (err error)
	GetNoticeStrategyInfo(id int) (data *vo.NoticeStrategyGetInfoRes, err error)
	GetNoticeStrategyList(filter map[string]interface{}) (total int64, data []*vo.NoticeStrategyGetInfoRes, err error)
}

var (
	noticeStrategyServiceOnce sync.Once
	noticeStrategyService     NoticeStrategyEntityInterface
)

type noticeStrategyEntity struct {
	noticeStrategyRepo repository.NoticeStrategyRepoInterface // 依赖抽象
}

var _ NoticeStrategyEntityInterface = &noticeStrategyEntity{}

func NewNoticeStrategyEntity() NoticeStrategyEntityInterface {
	noticeStrategyServiceOnce.Do(func() {
		noticeStrategyService = &noticeStrategyEntity{
			noticeStrategyRepo: persistence.NewNoticeStrategyRepo(),
		}
	})
	return noticeStrategyService
}

func (service *noticeStrategyEntity) AddNoticeStrategy(data *vo.NoticeStrategyCreateReq) (id int, err error) {
	// 参数处理
	confInfo := &model.NoticeStrategy{}
	//confInfo.Name = data.Name
	//confInfo.Gender = *data.Gender
	return service.noticeStrategyRepo.Create(confInfo)
}

func (service *noticeStrategyEntity) ModNoticeStrategy(id int, data map[string]interface{}) (err error) {
	return service.noticeStrategyRepo.Update(id, data)
}

func (service *noticeStrategyEntity) DelNoticeStrategy(id int) (err error) {
	return service.noticeStrategyRepo.Delete(id)
}

func (service *noticeStrategyEntity) GetNoticeStrategyInfo(id int) (data *vo.NoticeStrategyGetInfoRes, err error) {
	confInfo, err := service.noticeStrategyRepo.GetInfo(id)
	// 响应处理
	data = &vo.NoticeStrategyGetInfoRes{}
	data.Id = confInfo.Id
	//data.Name = confInfo.Name
	//data.Gender = confInfo.Gender
	data.SetEnabledStateDisplayName()
	return
}

func (service *noticeStrategyEntity) GetNoticeStrategyList(filter map[string]interface{}) (total int64, data []*vo.NoticeStrategyGetInfoRes, err error) {
	total, confList, err := service.noticeStrategyRepo.GetList(filter)
	// 响应处理
	data = make([]*vo.NoticeStrategyGetInfoRes, 0)
	for _, conf := range confList {
		item := &vo.NoticeStrategyGetInfoRes{}
		item.Id = conf.Id
		//item.Name = conf.Name
		//item.Gender = conf.Gender
		item.SetEnabledStateDisplayName()
		data = append(data, item)
	}
	return
}