package repository

import (
	"go_project/infra/model"
)

//go:generate mockgen -source=./notice_conf.go -destination ./mock/mock_notice_conf.go -package mock
type NoticeConfRepoInterface interface {
	Create(data *model.NoticeConf) (id int, err error)
	Update(id int, data map[string]interface{}) (err error)
	Delete(id int) (err error)
	GetInfo(id int) (data *model.NoticeConf, err error)
	GetList(filter map[string]interface{}) (total int64, data []*model.NoticeConf, err error)
}
