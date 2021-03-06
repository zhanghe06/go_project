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
	repoOnce sync.Once
	repo     repository.UserRepoInterface
)

type UserRepository struct {
	db *gorm.DB
	// TODO ETCD
}

var _ repository.UserRepoInterface = &UserRepository{}

func NewUserRepo() repository.UserRepoInterface {
	repoOnce.Do(func() {
		repo = &UserRepository{
			db: db.NewDB(),
		}
	})
	return repo
}

func (repo *UserRepository) Create(data *model.User) (id int, err error) {
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

func (repo *UserRepository) Update(id int, data map[string]interface{}) (err error) {
	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id
	condition["deleted_state"] = enums.NotDeleted

	// 更新时间
	currentTime := time.Now()
	data["updated_at"] = currentTime

	err = repo.db.Model(&model.User{}).Where(condition).Updates(data).Error
	return
}

func (repo *UserRepository) Delete(id int) (err error) {
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

	err = repo.db.Model(&model.User{}).Where(condition).Updates(data).Error
	return
}

func (repo *UserRepository) GetInfo(id int) (data *model.User, err error) {
	// 临时打印SQL
	// err = repo.db.Debug().First(&data, id).Error

	// 条件处理
	condition := make(map[string]interface{})
	condition["id"] = id
	condition["deleted_state"] = enums.NotDeleted

	err = repo.db.Where(condition).First(&data).Error
	return
}

func (repo *UserRepository) GetList(filter map[string]interface{}, args ...interface{}) (total int64, data []*model.User, err error) {
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
	dbQuery := repo.db.Model(&model.User{}).Where(condition)
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
