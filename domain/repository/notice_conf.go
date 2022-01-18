package repository

import (
	"sap_cert_mgt/infra/model"
)

//go:generate mockgen -source=./notice_conf.go -destination ./mock/mock_notice_conf.go -package mock
type NoticeConfRepoInterface interface {
	GetEmail() (data *model.NoticeConf, err error)
	ModEmail(data map[string]interface{}, updatedBy string) (err error)
}
