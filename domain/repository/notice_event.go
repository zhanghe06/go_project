package repository

import (
	"sap_cert_mgt/infra/model"
)

//go:generate mockgen -source=./notice_event.go -destination ./mock/mock_notice_event.go -package mock
type NoticeEventRepoInterface interface {
	Create(data *model.NoticeEvent) (id int, err error)
	Update(id int, data map[string]interface{}) (err error)
	Delete(id int) (err error)
	GetInfo(id int) (data *model.NoticeEvent, err error)
	GetList(filter map[string]interface{}, args ...interface{}) (total int64, data []*model.NoticeEvent, err error)
}
