package entity

import (
	"sap_cert_mgt/domain/repository"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/model"
	"sap_cert_mgt/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./operation_log.go -destination ./mock/mock_operation_log.go -package mock
type OperationLogEntityInterface interface {
	AddOperationLog(data *vo.OperationLogCreateReq, createdBy string) (id int, err error)
	DelOperationLog(id int, deletedBy string) (err error)
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
	logInfo.OpType = data.OpType
	logInfo.RsType = data.RsType
	logInfo.RsId = data.RsId
	logInfo.OpDetail = data.OpDetail
	logInfo.OpError = data.OpError
	return service.operationLogRepo.Create(logInfo, createdBy)
}

func (service *operationLogEntity) DelOperationLog(id int, deletedBy string) (err error) {
	return service.operationLogRepo.Delete(id, deletedBy)
}

func (service *operationLogEntity) GetOperationLogInfo(id int) (data *vo.OperationLogGetInfoRes, err error) {
	resInfo, err := service.operationLogRepo.GetInfo(id)
	// 响应处理
	data = &vo.OperationLogGetInfoRes{}
	data.Id = resInfo.Id
	data.OpType = resInfo.OpType
	data.RsType = resInfo.RsType
	data.RsId = resInfo.RsId
	data.OpDetail = resInfo.OpDetail
	data.OpError = resInfo.OpError
	return
}

func (service *operationLogEntity) GetOperationLogList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.OperationLogGetInfoRes, err error) {
	total, resList, err := service.operationLogRepo.GetList(filter, args...)
	// 响应处理
	data = make([]*vo.OperationLogGetInfoRes, 0)
	for _, resInfo := range resList {
		item := &vo.OperationLogGetInfoRes{}
		item.Id = resInfo.Id
		item.OpType = resInfo.OpType
		item.RsType = resInfo.RsType
		item.RsId = resInfo.RsId
		item.OpDetail = resInfo.OpDetail
		item.OpError = resInfo.OpError
		data = append(data, item)
	}
	return
}
