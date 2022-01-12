package repository

import (
	"go_project/infra/model"
)

//go:generate mockgen -source=./cert.go -destination ./mock/mock_cert.go -package mock
type CertRepoInterface interface {
	Create(data *model.Cert) (id int, err error)
	Update(id int, data map[string]interface{}) (err error)
	Delete(id int) (err error)
	GetInfo(id int) (data *model.Cert, err error)
	GetList(filter map[string]interface{}) (total int64, data []*model.Cert, err error)
}
