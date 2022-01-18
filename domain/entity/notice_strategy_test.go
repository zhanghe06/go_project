package entity

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	. "github.com/smartystreets/goconvey/convey"
	. "sap_cert_mgt/domain/repository/mock"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/model"
	"testing"
)

// go test -cover ./...
func TestNoticeStrategyEntity(t *testing.T) {
	ctrl := gomock.NewController(t) // 初始化 controller
	defer ctrl.Finish()

	noticeStrategyRepo := NewMockNoticeStrategyRepoInterface(ctrl) // 初始化 mock
	opLogRepo := NewMockOperationLogRepoInterface(ctrl)            // 初始化 mock

	noticeStrategyInst := noticeStrategyEntity{
		noticeStrategyRepo: noticeStrategyRepo,
		opLogRepo:          opLogRepo,
	}

	Convey("Convey Test Notice Strategy Add Info Entity", t, func() {
		var id int
		var err error
		var instReq *vo.NoticeStrategyCreateReq
		var createdBy string

		Convey("Add Success", func() {
			id = 1
			err = nil
			createdBy = ""
			noticeType := new(int)
			*noticeType = 0
			triggerThreshold := new(int)
			*triggerThreshold = 0
			enabledState := new(int)
			*enabledState = 0

			instReq = &vo.NoticeStrategyCreateReq{
				NoticeType:       noticeType,
				TriggerThreshold: triggerThreshold,
				EnabledState:     enabledState,
			}
			noticeStrategyRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(id, err)
			opLogRepo.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes().Return(id, err)
			idRes, errRes := noticeStrategyInst.AddNoticeStrategy(instReq, createdBy)
			assert.Equal(t, idRes, id)
			assert.Equal(t, errRes, err)
		})
	})

	Convey("Convey Test Notice Strategy Get Info Entity", t, func() {
		var id int
		var err error
		//var data *vo.NoticeStrategyGetInfoRes

		Convey("Get Success", func() {
			id = 1
			err = nil
			repoDataRes := &model.NoticeStrategy{
				Id: id,
			}
			noticeStrategyRepo.EXPECT().GetInfo(id).Return(repoDataRes, err)
			dataRes, errRes := noticeStrategyInst.GetNoticeStrategyInfo(id)
			assert.Equal(t, dataRes.Id, id)
			assert.Equal(t, errRes, err)
		})
	})

	Convey("Convey Test Notice Strategy Mod Info Entity", t, func() {
		var id int
		var updatedBy string
		var err error
		var data map[string]interface{}

		Convey("Mod Success", func() {
			err = nil
			data = make(map[string]interface{})
			id = 1
			updatedBy = ""

			noticeStrategyRepo.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(err)
			opLogRepo.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes().Return(id, err)
			errRes := noticeStrategyInst.ModNoticeStrategy(id, data, updatedBy)
			assert.Equal(t, errRes, err)
		})
		Convey("Mod Failure", func() {
			err = errors.New("")
			data = make(map[string]interface{})
			id = 1
			updatedBy = ""

			noticeStrategyRepo.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(err)
			opLogRepo.EXPECT().Create(gomock.Any(), gomock.Any()).AnyTimes().Return(id, err)
			errRes := noticeStrategyInst.ModNoticeStrategy(id, data, updatedBy)
			assert.Equal(t, errRes, err)
		})
	})

	Convey("Convey Test Notice Strategy List Info Entity", t, func() {
		var repoTotalRes int64
		var err error
		var filter map[string]interface{}
		var repoDataRes []*model.NoticeStrategy

		Convey("List Success", func() {
			err = nil
			filter = make(map[string]interface{})
			repoTotalRes = 0

			noticeStrategyRepo.EXPECT().GetList(filter).Return(repoTotalRes, repoDataRes, err)
			totalRes, _, errRes := noticeStrategyInst.GetNoticeStrategyList(filter)
			assert.Equal(t, totalRes, repoTotalRes)
			assert.Equal(t, errRes, err)
		})
	})

	Convey("Convey Test Notice Strategy Delete Info Entity", t, func() {
		var id int
		var deletedBy string
		var err error

		Convey("Delete Success", func() {
			id = 1
			deletedBy = ""
			err = nil

			noticeStrategyRepo.EXPECT().Delete(id, deletedBy).Return(err)
			opLogRepo.EXPECT().Create(gomock.Any(), deletedBy).AnyTimes().Return(id, err)
			errRes := noticeStrategyInst.DelNoticeStrategy(id, deletedBy)
			assert.Equal(t, errRes, err)
		})
	})
}
