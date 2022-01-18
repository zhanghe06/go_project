package entity

import (
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	. "github.com/smartystreets/goconvey/convey"
	. "sap_cert_mgt/domain/repository/mock"
	"sap_cert_mgt/domain/vo"
	"testing"
	"time"
)

// go test -cover ./...
func TestCertEntity(t *testing.T) {
	ctrl := gomock.NewController(t) // 初始化 controller
	defer ctrl.Finish()

	certRepo := NewMockCertRepoInterface(ctrl)          // 初始化 mock
	opLogRepo := NewMockOperationLogRepoInterface(ctrl) // 初始化 mock

	certInst := certEntity{
		certRepo:  certRepo,
		opLogRepo: opLogRepo,
	}

	Convey("Convey Test Add Cert Entity", t, func() {
		var err error
		var data *vo.CertCreateReq
		var createdBy string

		Convey("Add Cert Success", func() {
			err = nil
			version := new(int)
			*version = 0
			enabledState := new(int)
			*enabledState = 0
			notBeforeStr := "2021-01-01 00:00:00"
			notAfterStr := "2022-12-31 59:59:59"
			notBefore, _ := time.ParseInLocation("2006-01-02 15:04:05", notBeforeStr, time.Local)
			notAfter, _ := time.ParseInLocation("2006-01-02 15:04:05", notAfterStr, time.Local)
			data = &vo.CertCreateReq{
				AuthId:             "",        // 客户端ID
				PVersion:           "",        // 接口版本
				ContRep:            "",        // 内容存储库
				SerialNumber:       "",        // 证书序列号
				Version:            version,   // 证书版本（0:V1,1:V2,2:V3）
				IssuerName:         "",        // 颁发机构
				SignatureAlgorithm: "",        // 签名算法
				NotBefore:          notBefore, // 有效期开始时间
				NotAfter:           notAfter,  // 有效期结束时间
				EnabledState:       enabledState,
			}
			createdBy = ""
			id := 1
			certRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(id, err)
			certRepo.EXPECT().Enable(id, createdBy).Return(err)
			opLogRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(id, err)
			idRes, errRes := certInst.AddCert(data, createdBy)
			assert.Equal(t, idRes, id)
			assert.Equal(t, errRes, err)
		})
	})

	Convey("Convey Test Del Cert Entity", t, func() {
		var err error
		var deletedBy string

		Convey("Del Cert Success", func() {
			err = nil
			deletedBy = ""
			id := 1
			certRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(err)
			opLogRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(id, err)
			errRes := certInst.DelCert(id, deletedBy)
			assert.Equal(t, errRes, err)
		})
	})
}
