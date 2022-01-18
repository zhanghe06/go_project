package entity

import (
	"sap_cert_mgt/domain/repository"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/model"
	"sap_cert_mgt/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./cert.go -destination ./mock/mock_cert.go -package mock
type CertEntityInterface interface {
	AddCert(data *vo.CertCreateReq, createdBy string) (id int, err error)
	DelCert(id int, deletedBy string) (err error)
	GetCertInfo(id int) (data *vo.CertGetInfoRes, err error)
	GetCertList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.CertGetInfoRes, err error)
}

var (
	certServiceOnce sync.Once
	certService     CertEntityInterface
)

type certEntity struct {
	certRepo  repository.CertRepoInterface         // 依赖抽象
	opLogRepo repository.OperationLogRepoInterface // 依赖抽象
}

var _ CertEntityInterface = &certEntity{}

func NewCertEntity() CertEntityInterface {
	certServiceOnce.Do(func() {
		certService = &certEntity{
			certRepo:  persistence.NewCertRepo(),
			opLogRepo: persistence.NewOperationLogRepo(),
		}
	})
	return certService
}

func (service *certEntity) AddCert(data *vo.CertCreateReq, createdBy string) (id int, err error) {
	// 参数处理
	certInfo := &model.Cert{}
	certInfo.AuthId = data.AuthId
	certInfo.PVersion = data.PVersion
	certInfo.ContRep = data.ContRep
	certInfo.SerialNumber = data.SerialNumber
	certInfo.Version = *data.Version
	certInfo.IssuerName = data.IssuerName
	certInfo.SignatureAlgorithm = data.SignatureAlgorithm
	certInfo.NotBefore = data.NotBefore
	certInfo.NotAfter = data.NotAfter
	certInfo.EnabledState = *data.EnabledState
	// 创建证书
	id, err = service.certRepo.Create(certInfo, createdBy)
	if err != nil {
		return
	}
	// 启用证书
	err = service.certRepo.Enable(id, createdBy)
	if err != nil {
		return
	}
	// 操作日志
	opLogData := &model.OperationLog{
		OpType:   "create",
		RsType:   "cert",
		RsId:     id,
		OpDetail: "",
		OpError:  "",
	}
	_, err = service.opLogRepo.Create(opLogData, createdBy)
	return
}

func (service *certEntity) DelCert(id int, deletedBy string) (err error) {
	err = service.certRepo.Delete(id, deletedBy)
	if err != nil {
		return
	}
	// 操作日志
	opLogData := &model.OperationLog{
		OpType:   "delete",
		RsType:   "cert",
		RsId:     id,
		OpDetail: "",
		OpError:  "",
	}
	_, err = service.opLogRepo.Create(opLogData, deletedBy)
	return
}

func (service *certEntity) GetCertInfo(id int) (data *vo.CertGetInfoRes, err error) {
	certInfo, err := service.certRepo.GetInfo(id)
	// 响应处理
	data = &vo.CertGetInfoRes{}
	data.Id = certInfo.Id
	data.AuthId = certInfo.AuthId
	data.PVersion = certInfo.PVersion
	data.ContRep = certInfo.ContRep
	data.SerialNumber = certInfo.SerialNumber
	data.Version = certInfo.Version
	data.IssuerName = certInfo.IssuerName
	data.SignatureAlgorithm = certInfo.SignatureAlgorithm
	data.NotBefore = certInfo.NotBefore
	data.NotAfter = certInfo.NotAfter
	data.EnabledState = certInfo.EnabledState
	data.SetVersionDisplayName()
	data.SetEnabledStateDisplayName()
	return
}

func (service *certEntity) GetCertList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.CertGetInfoRes, err error) {
	total, certList, err := service.certRepo.GetList(filter, args...)
	// 响应处理
	data = make([]*vo.CertGetInfoRes, 0)
	for _, certInfo := range certList {
		item := &vo.CertGetInfoRes{}
		item.Id = certInfo.Id
		item.AuthId = certInfo.AuthId
		item.PVersion = certInfo.PVersion
		item.ContRep = certInfo.ContRep
		item.SerialNumber = certInfo.SerialNumber
		item.Version = certInfo.Version
		item.IssuerName = certInfo.IssuerName
		item.SignatureAlgorithm = certInfo.SignatureAlgorithm
		item.NotBefore = certInfo.NotBefore
		item.NotAfter = certInfo.NotAfter
		item.EnabledState = certInfo.EnabledState
		item.SetVersionDisplayName()
		item.SetEnabledStateDisplayName()
		data = append(data, item)
	}
	return
}
