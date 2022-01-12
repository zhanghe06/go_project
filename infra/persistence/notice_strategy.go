package persistence

import (
	"go_project/domain/enums"
	"go_project/domain/repository"
	"go_project/infra/db"
	"go_project/infra/model"
	"gorm.io/gorm"
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

func (repo *NoticeStrategyRepository) Create(data *model.NoticeStrategy) (id int, err error) {
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

func (repo *NoticeStrategyRepository) Update(id int, data map[string]interface{}) (err error) {
	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id
	condition["deleted_state"] = enums.NotDeleted

	// 更新时间
	currentTime := time.Now()
	data["updated_at"] = currentTime

	err = repo.db.Model(&model.NoticeStrategy{}).Where(condition).Updates(data).Error
	return
}

func (repo *NoticeStrategyRepository) Delete(id int) (err error) {
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

func (repo *NoticeStrategyRepository) GetList(filter map[string]interface{}) (total int64, data []*model.NoticeStrategy, err error) {
	// 条件处理
	limit := 10
	offset := 0
	condition := make(map[string]interface{})
	for k, v := range filter {
		if k == "limit" {
			limit = v.(int)
		} else if k == "offset" {
			offset = v.(int)
		} else {
			condition[k] = v
		}
	}
	condition["deleted_state"] = enums.NotDeleted

	// 总记录数
	userObj := repo.db.Model(&model.NoticeStrategy{}).Where(condition)
	userObj.Count(&total)

	// 分页查询
	err = userObj.Limit(limit).Offset(offset).Find(&data).Error
	return
}