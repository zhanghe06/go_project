package repository

import (
	"go_project/infra/model"
)

//go:generate mockgen -source=./notice_strategy.go -destination ./mock/mock_notice_strategy.go -package mock
type NoticeStrategyRepoInterface interface {
	Create(data *model.NoticeStrategy, createdBy string) (id int, err error)
	Update(id int, data map[string]interface{}, updatedBy string) (err error)
	Delete(id int, deletedBy string) (err error)
	GetInfo(id int) (data *model.NoticeStrategy, err error)
	GetList(filter map[string]interface{}, args ...interface{}) (total int64, data []*model.NoticeStrategy, err error)
}
