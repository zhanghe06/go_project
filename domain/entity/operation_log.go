package entity

import (
	"go_project/domain/repository"
	"go_project/domain/vo"
	"go_project/infra/model"
	"go_project/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./operation_log.go -destination ./mock/mock_operation_log.go -package mock
type OperationLogEntityInterface interface {
	AddOperationLog(data *vo.OperationLogCreateReq, createdBy string) (id int, err error)
	DelOperationLog(id int, deletedBy string) (err error)
	ModOperationLog(id int, data map[string]interface{}, updatedBy string) (err error)
	GetOperationLogInfo(id int) (data *vo.OperationLogGetInfoRes, err error)
	GetOperationLogList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.OperationLogGetInfoRes, err error)
}

var (
	operationLogServiceOnce sync.Once
	operationLogService     OperationLogEntityInterface
)

type operationLogEntity struct {
	operationLogRepo repository.OperationLogRepoInterface // 依赖抽象
}

var _ OperationLogEntityInterface = &operationLogEntity{}

func NewOperationLogEntity() OperationLogEntityInterface {
	operationLogServiceOnce.Do(func() {
		operationLogService = &operationLogEntity{
			operationLogRepo: persistence.NewOperationLogRepo(),
		}
	})
	return operationLogService
}

func (service *operationLogEntity) AddOperationLog(data *vo.OperationLogCreateReq, createdBy string) (id int, err error) {
	// 参数处理
	logInfo := &model.OperationLog{}
	//logInfo.Name = data.Name
	//logInfo.Gender = *data.Gender
	return service.operationLogRepo.Create(logInfo, createdBy)
}

func (service *operationLogEntity) ModOperationLog(id int, data map[string]interface{}, updatedBy string) (err error) {
	return service.operationLogRepo.Update(id, data, updatedBy)
}

func (service *operationLogEntity) DelOperationLog(id int, deletedBy string) (err error) {
	return service.operationLogRepo.Delete(id, deletedBy)
}

func (service *operationLogEntity) GetOperationLogInfo(id int) (data *vo.OperationLogGetInfoRes, err error) {
	resInfo, err := service.operationLogRepo.GetInfo(id)
	// 响应处理
	data = &vo.OperationLogGetInfoRes{}
	data.Id = resInfo.Id
	//data.Name = resInfo.Name
	//data.Gender = resInfo.Gender
	data.SetGenderDisplayName()
	return
}

func (service *operationLogEntity) GetOperationLogList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.OperationLogGetInfoRes, err error) {
	total, resList, err := service.operationLogRepo.GetList(filter, args...)
	// 响应处理
	data = make([]*vo.OperationLogGetInfoRes, 0)
	for _, resInfo := range resList {
		item := &vo.OperationLogGetInfoRes{}
		item.Id = resInfo.Id
		//item.Name = resInfo.Name
		//item.Gender = resInfo.Gender
		item.SetGenderDisplayName()
		data = append(data, item)
	}
	return
}
