package entity

import (
	"go_project/domain/repository"
	"go_project/domain/vo"
	"go_project/infra/model"
	"go_project/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./notice_conf.go -destination ./mock/mock_notice_conf.go -package mock
type NoticeConfEntityInterface interface {
	AddNoticeConf(data *vo.NoticeConfCreateReq) (id int, err error)
	DelNoticeConf(id int) (err error)
	ModNoticeConf(id int, data map[string]interface{}) (err error)
	GetNoticeConfInfo(id int) (data *vo.NoticeConfGetInfoRes, err error)
	GetNoticeConfList(filter map[string]interface{}) (total int64, data []*vo.NoticeConfGetInfoRes, err error)
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

func (service *noticeConfEntity) AddNoticeConf(data *vo.NoticeConfCreateReq) (id int, err error) {
	// 参数处理
	confInfo := &model.NoticeConf{}
	//confInfo.Name = data.Name
	//confInfo.Gender = *data.Gender
	return service.noticeConfRepo.Create(confInfo)
}

func (service *noticeConfEntity) ModNoticeConf(id int, data map[string]interface{}) (err error) {
	return service.noticeConfRepo.Update(id, data)
}

func (service *noticeConfEntity) DelNoticeConf(id int) (err error) {
	return service.noticeConfRepo.Delete(id)
}

func (service *noticeConfEntity) GetNoticeConfInfo(id int) (data *vo.NoticeConfGetInfoRes, err error) {
	confInfo, err := service.noticeConfRepo.GetInfo(id)
	// 响应处理
	data = &vo.NoticeConfGetInfoRes{}
	data.Id = confInfo.Id
	//data.Name = confInfo.Name
	//data.Gender = confInfo.Gender
	data.SetGenderDisplayName()
	return
}

func (service *noticeConfEntity) GetNoticeConfList(filter map[string]interface{}) (total int64, data []*vo.NoticeConfGetInfoRes, err error) {
	total, confList, err := service.noticeConfRepo.GetList(filter)
	// 响应处理
	data = make([]*vo.NoticeConfGetInfoRes, 0)
	for _, conf := range confList {
		item := &vo.NoticeConfGetInfoRes{}
		item.Id = conf.Id
		//item.Name = conf.Name
		//item.Gender = conf.Gender
		item.SetGenderDisplayName()
		data = append(data, item)
	}
	return
}
