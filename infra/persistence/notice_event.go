package persistence

import (
	"gorm.io/gorm"
	"sap_cert_mgt/domain/enums"
	"sap_cert_mgt/domain/repository"
	"sap_cert_mgt/infra/db"
	"sap_cert_mgt/infra/model"
	"sync"
	"time"
)

var (
	noticeEventRepoOnce sync.Once
	noticeEventRepo     repository.NoticeEventRepoInterface
)

type NoticeEventRepository struct {
	db *gorm.DB
	// TODO ETCD
}

var _ repository.NoticeEventRepoInterface = &NoticeEventRepository{}

func NewNoticeEventRepo() repository.NoticeEventRepoInterface {
	noticeEventRepoOnce.Do(func() {
		noticeEventRepo = &NoticeEventRepository{
			db: db.NewDB(),
		}
	})
	return noticeEventRepo
}

func (repo *NoticeEventRepository) Create(data *model.NoticeEvent) (id int, err error) {
	// 创建时间
	currentTime := time.Now()
	data.CreatedAt = currentTime

	result := repo.db.Create(&data)
	id = data.Id
	err = result.Error
	//data.Id             // 返回插入数据的主键
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数
	return
}

func (repo *NoticeEventRepository) Update(id int, data map[string]interface{}) (err error) {
	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id
	condition["deleted_state"] = enums.NotDeleted

	// 更新时间
	currentTime := time.Now()
	data["updated_at"] = currentTime

	err = repo.db.Model(&model.NoticeEvent{}).Where(condition).Updates(data).Error
	return
}

func (repo *NoticeEventRepository) Delete(id int) (err error) {
	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id
	condition["deleted_state"] = enums.NotDeleted

	// 逻辑删除
	data := make(map[string]interface{})
	data["deleted_state"] = enums.HasDeleted
	// 删除时间
	currentTime := time.Now()
	data["deleted_at"] = currentTime

	err = repo.db.Model(&model.NoticeEvent{}).Where(condition).Updates(data).Error
	return
}

func (repo *NoticeEventRepository) GetInfo(id int) (data *model.NoticeEvent, err error) {
	// 临时打印SQL
	// err = repo.db.Debug().First(&data, id).Error

	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id
	condition["deleted_state"] = enums.NotDeleted

	err = repo.db.Where(condition).First(&data).Error
	return
}

func (repo *NoticeEventRepository) GetList(filter map[string]interface{}, args ...interface{}) (total int64, data []*model.NoticeEvent, err error) {
	// 条件处理
	limit := 10
	offset := 0
	condition := make(map[string]interface{})
	for k, v := range filter {
		if k == "limit" {
			limit = int(v.(float64))
		} else if k == "offset" {
			offset = int(v.(float64))
		} else {
			condition[k] = v
		}
	}
	condition["deleted_state"] = enums.NotDeleted

	// 总记录数
	dbQuery := repo.db.Model(&model.NoticeEvent{}).Where(condition)
	if len(args) >= 2 {
		dbQuery = dbQuery.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		dbQuery = dbQuery.Where(args[0])
	}
	dbQuery.Count(&total)

	// 分页查询
	err = dbQuery.Limit(limit).Offset(offset).Find(&data).Error
	return
}
