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
	noticeStrategyRepoOnce sync.Once
	noticeStrategyRepo     repository.NoticeStrategyRepoInterface
)

type NoticeStrategyRepository struct {
	db *gorm.DB
	// TODO ETCD
}

var _ repository.NoticeStrategyRepoInterface = &NoticeStrategyRepository{}

func NewNoticeStrategyRepo() repository.NoticeStrategyRepoInterface {
	noticeStrategyRepoOnce.Do(func() {
		noticeStrategyRepo = &NoticeStrategyRepository{
			db: db.NewDB(),
		}
	})
	return noticeStrategyRepo
}

func (repo *NoticeStrategyRepository) Create(data *model.NoticeStrategy, createdBy string) (id int, err error) {
	// 创建时间
	currentTime := time.Now()
	data.CreatedAt = currentTime
	data.CreatedBy = createdBy

	result := repo.db.Create(&data)
	id = data.Id
	err = result.Error
	//data.Id             // 返回插入数据的主键
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数
	return
}

func (repo *NoticeStrategyRepository) Update(id int, data map[string]interface{}, updatedBy string) (err error) {
	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id
	condition["deleted_state"] = enums.NotDeleted

	// 更新时间
	currentTime := time.Now()
	data["updated_at"] = currentTime
	data["updated_by"] = updatedBy

	err = repo.db.Model(&model.NoticeStrategy{}).Where(condition).Updates(data).Error
	return
}

func (repo *NoticeStrategyRepository) Delete(id int, deletedBy string) (err error) {
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
	data["deleted_by"] = deletedBy

	err = repo.db.Model(&model.NoticeStrategy{}).Where(condition).Updates(data).Error
	return
}

func (repo *NoticeStrategyRepository) GetInfo(id int) (data *model.NoticeStrategy, err error) {
	// 临时打印SQL
	// err = repo.db.Debug().First(&data, id).Error

	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id
	condition["deleted_state"] = enums.NotDeleted

	err = repo.db.Where(condition).First(&data).Error
	return
}

func (repo *NoticeStrategyRepository) GetList(filter map[string]interface{}, args ...interface{}) (total int64, data []*model.NoticeStrategy, err error) {
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
	dbQuery := repo.db.Model(&model.NoticeStrategy{}).Where(condition)
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
