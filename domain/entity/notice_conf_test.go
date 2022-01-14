package entity

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	. "github.com/smartystreets/goconvey/convey"
	. "go_project/domain/repository/mock"
	"go_project/domain/vo"
	"go_project/infra/model"
	"testing"
)

// go test -cover ./...
func TestNoticeConfEntity(t *testing.T) {
	ctrl := gomock.NewController(t) // 初始化 controller
	defer ctrl.Finish()

	noticeConfRepo := NewMockNoticeConfRepoInterface(ctrl) // 初始化 mock

	noticeConfInst := noticeConfEntity{
		noticeConfRepo: noticeConfRepo,
	}

	Convey("Convey Test Notice Conf Get Email Entity", t, func() {
		var err error
		var data *vo.NoticeConfGetEmailRes

		Convey("Get Success", func() {
			err = nil
			data = &vo.NoticeConfGetEmailRes{}
			repoDataRes := &model.NoticeConf{}
			noticeConfRepo.EXPECT().GetEmail().Return(repoDataRes, err)
			dataRes, errRes := noticeConfInst.GetNoticeConfEmail()
			assert.Equal(t, dataRes, data)
			assert.Equal(t, errRes, err)
		})
	})

	Convey("Convey Test Notice Conf Mod Email Entity", t, func() {
		var updatedBy string
		var err error
		var data map[string]interface{}

		Convey("Mod Success", func() {
			err = nil
			data = make(map[string]interface{})
			data["server_host"] = ""
			data["server_port"] = ""
			data["from_name"] = ""
			data["from_email"] = ""
			data["from_passwd"] = ""
			updatedBy = ""

			noticeConfRepo.EXPECT().ModEmail(gomock.Any(), gomock.Any()).Return(err)
			errRes := noticeConfInst.ModNoticeConfEmail(data, updatedBy)
			assert.Equal(t, errRes, err)
		})
		Convey("Mod Failure", func() {
			err = errors.New("")
			data = make(map[string]interface{})
			data["server_host"] = ""
			data["server_port"] = ""
			data["from_name"] = ""
			data["from_email"] = ""
			data["from_passwd"] = ""
			updatedBy = ""

			noticeConfRepo.EXPECT().ModEmail(gomock.Any(), gomock.Any()).Return(err)
			errRes := noticeConfInst.ModNoticeConfEmail(data, updatedBy)
			assert.Equal(t, errRes, err)
		})
	})
}
