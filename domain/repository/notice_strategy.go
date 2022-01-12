package repository

import (
	"go_project/infra/model"
)

//go:generate mockgen -source=./notice_strategy.go -destination ./mock/mock_notice_strategy.go -package mock
type NoticeStrategyRepoInterface interface {
	Create(data *model.NoticeStrategy) (id int, err error)
	Update(id int, data map[string]interface{}) (err error)
	Delete(id int) (err error)
	GetInfo(id int) (data *model.NoticeStrategy, err error)
	GetList(filter map[string]interface{}) (total int64, data []*model.NoticeStrategy, err error)
}
