package entity

import (
	"go_project/domain/repository"
	"go_project/domain/vo"
	"go_project/infra/model"
	"go_project/infra/persistence"
	"sync"
)

//go:generate mockgen -source=./cert.go -destination ./mock/mock_cert.go -package mock
type CertEntityInterface interface {
	AddCert(data *vo.CertCreateReq, createdBy string) (id int, err error)
	DelCert(id int, deletedBy string) (err error)
	ModCert(id int, data map[string]interface{}, updatedBy string) (err error)
	GetCertInfo(id int) (data *vo.CertGetInfoRes, err error)
	GetCertList(filter map[string]interface{}, args ...interface{}) (total int64, data []*vo.CertGetInfoRes, err error)
}

var (
	certServiceOnce sync.Once
	certService     CertEntityInterface
)

type certEntity struct {
	certRepo repository.CertRepoInterface // 依赖抽象
}

var _ CertEntityInterface = &certEntity{}

func NewCertEntity() CertEntityInterface {
	certServiceOnce.Do(func() {
		certService = &certEntity{
			certRepo: persistence.NewCertRepo(),
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
	return
}

func (service *certEntity) ModCert(id int, data map[string]interface{}, updatedBy string) (err error) {
	return service.certRepo.Update(id, data, updatedBy)
}

func (service *certEntity) DelCert(id int, deletedBy string) (err error) {
	return service.certRepo.Delete(id, deletedBy)
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
