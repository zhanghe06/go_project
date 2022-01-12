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
	AddCert(data *vo.CertCreateReq) (id int, err error)
	DelCert(id int) (err error)
	ModCert(id int, data map[string]interface{}) (err error)
	GetCertInfo(id int) (data *vo.CertGetInfoRes, err error)
	GetCertList(filter map[string]interface{}) (total int64, data []*vo.CertGetInfoRes, err error)
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

func (service *certEntity) AddCert(data *vo.CertCreateReq) (id int, err error) {
	// 参数处理
	certInfo := &model.Cert{}
	// todo
	//certInfo.Name = data.Name
	//certInfo.Gender = *data.Gender
	return service.certRepo.Create(certInfo)
}

func (service *certEntity) ModCert(id int, data map[string]interface{}) (err error) {
	return service.certRepo.Update(id, data)
}

func (service *certEntity) DelCert(id int) (err error) {
	return service.certRepo.Delete(id)
}

func (service *certEntity) GetCertInfo(id int) (data *vo.CertGetInfoRes, err error) {
	certInfo, err := service.certRepo.GetInfo(id)
	// 响应处理
	data = &vo.CertGetInfoRes{}
	data.Id = certInfo.Id
	//data.Name = certInfo.Name
	//data.Gender = certInfo.Gender
	data.SetGenderDisplayName()
	return
}

func (service *certEntity) GetCertList(filter map[string]interface{}) (total int64, data []*vo.CertGetInfoRes, err error) {
	total, certList, err := service.certRepo.GetList(filter)
	// 响应处理
	data = make([]*vo.CertGetInfoRes, 0)
	for _, cert := range certList {
		item := &vo.CertGetInfoRes{}
		item.Id = cert.Id
		//item.Name = cert.Name
		//item.Gender = cert.Gender
		item.SetGenderDisplayName()
		data = append(data, item)
	}
	return
}
