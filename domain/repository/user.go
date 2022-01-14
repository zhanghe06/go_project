package repository

import (
	"go_project/infra/model"
)

//go:generate mockgen -source=./user.go -destination ./mock/mock_user.go -package mock
type UserRepoInterface interface {
	Create(data *model.User) (id int, err error)
	Update(id int, data map[string]interface{}) (err error)
	Delete(id int) (err error)
	GetInfo(id int) (data *model.User, err error)
	GetList(filter map[string]interface{}, args ...interface{}) (total int64, data []*model.User, err error)
}
