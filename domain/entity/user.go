package entity

import (
	"go_project/domain/repository"
	"go_project/domain/vo"
	"go_project/infra/model"
	"go_project/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./user.go -destination ./mock/mock_user.go -package mock
type UserEntityInterface interface {
	AddUser(data *vo.UserCreateReq) (id int, err error)
	DelUser(id int) (err error)
	ModUser(id int, data map[string]interface{}) (err error)
	GetUserInfo(id int) (data *vo.UserGetInfoRes, err error)
	GetUserList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.UserGetInfoRes, err error)
}

var (
	entityOnce sync.Once
	entity     UserEntityInterface
)

type userEntity struct {
	userRepo repository.UserRepoInterface // 依赖抽象
}

var _ UserEntityInterface = &userEntity{}

func NewUserEntity() UserEntityInterface {
	entityOnce.Do(func() {
		entity = &userEntity{
			userRepo: persistence.NewUserRepo(),
		}
	})
	return entity
}

func (service *userEntity) AddUser(data *vo.UserCreateReq) (id int, err error) {
	// 参数处理
	userInfo := &model.User{}
	userInfo.Name = data.Name
	userInfo.Gender = *data.Gender
	return service.userRepo.Create(userInfo)
}

func (service *userEntity) ModUser(id int, data map[string]interface{}) (err error) {
	return service.userRepo.Update(id, data)
}

func (service *userEntity) DelUser(id int) (err error) {
	return service.userRepo.Delete(id)
}

func (service *userEntity) GetUserInfo(id int) (data *vo.UserGetInfoRes, err error) {
	userInfo, err := service.userRepo.GetInfo(id)
	// 响应处理
	data = &vo.UserGetInfoRes{}
	data.Id = userInfo.Id
	data.Name = userInfo.Name
	data.Gender = userInfo.Gender
	data.SetGenderDisplayName()
	return
}

func (service *userEntity) GetUserList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.UserGetInfoRes, err error) {
	total, userList, err := service.userRepo.GetList(filter, args...)
	// 响应处理
	data = make([]*vo.UserGetInfoRes, 0)
	for _, user := range userList {
		item := &vo.UserGetInfoRes{}
		item.Id = user.Id
		item.Name = user.Name
		item.Gender = user.Gender
		item.SetGenderDisplayName()
		data = append(data, item)
	}
	return
}
