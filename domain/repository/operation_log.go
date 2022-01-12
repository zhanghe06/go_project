package repository

import (
	"go_project/infra/model"
)

//go:generate mockgen -source=./operation_log.go -destination ./mock/mock_operation_log.go -package mock
type OperationLogRepoInterface interface {
	Create(data *model.OperationLog, createdBy string) (id int, err error)
	Update(id int, data map[string]interface{}, updatedBy string) (err error)
	Delete(id int, deletedBy string) (err error)
	GetInfo(id int) (data *model.OperationLog, err error)
	GetList(filter map[string]interface{}) (total int64, data []*model.OperationLog, err error)
}
