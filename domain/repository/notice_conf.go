package repository

import (
	"go_project/infra/model"
)

//go:generate mockgen -source=./notice_conf.go -destination ./mock/mock_notice_conf.go -package mock
type NoticeConfRepoInterface interface {
	GetEmail() (data *model.NoticeConf, err error)
	ModEmail(data map[string]interface{}) (err error)
}
