package persistence

import (
	"encoding/json"
	"go_project/domain/enums"
	"go_project/domain/repository"
	"go_project/infra/db"
	"go_project/infra/model"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	noticeConfRepoOnce sync.Once
	noticeConfRepo     repository.NoticeConfRepoInterface
)

type NoticeConfRepository struct {
	db *gorm.DB
	// TODO ETCD
}

var _ repository.NoticeConfRepoInterface = &NoticeConfRepository{}

func NewNoticeConfRepo() repository.NoticeConfRepoInterface {
	noticeConfRepoOnce.Do(func() {
		noticeConfRepo = &NoticeConfRepository{
			db: db.NewDB(),
		}
	})
	return noticeConfRepo
}

func (repo *NoticeConfRepository) GetEmail() (data *model.NoticeConf, err error) {
	// 临时打印SQL
	// err = repo.db.Debug().First(&data, id).Error

	// 条件处理
	condition := make(map[string]interface{})
	condition["notice_type"] = enums.NoticeTypeEmail
	condition["deleted_state"] = enums.NotDeleted

	err = repo.db.Where(condition).First(&data).Error
	return
}

func (repo *NoticeConfRepository) ModEmail(data map[string]interface{}, updatedBy string) (err error) {
	// 临时打印SQL
	// err = repo.db.Debug().First(&data, id).Error

	// 参数处理
	configData, err := json.Marshal(data)
	if err != nil {
		return
	}

	// 条件处理
	var noticeConf *model.NoticeConf
	condition := make(map[string]interface{})
	condition["notice_type"] = enums.NoticeTypeEmail
	condition["deleted_state"] = enums.NotDeleted
	err = repo.db.Where(condition).First(&noticeConf).Error
	// 当时间
	currentTime := time.Now()

	if err != nil {
		// 创建记录
		noticeConf = &model.NoticeConf{}
		noticeConf.NoticeType = enums.NoticeTypeEmail
		noticeConf.ConfigData = string(configData)
		noticeConf.CreatedAt = currentTime
		noticeConf.CreatedBy = updatedBy
		result := repo.db.Create(&noticeConf)
		err = result.Error
	} else {
		// 更新记录
		noticeConf.ConfigData = string(configData)
		noticeConf.UpdatedAt = currentTime
		noticeConf.UpdatedBy = updatedBy
		repo.db.Save(&noticeConf)
	}
	return
}
