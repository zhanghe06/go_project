package entity

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	. "github.com/smartystreets/goconvey/convey"
	. "go_project/domain/repository/mock"
	"go_project/domain/vo"
	"testing"
)

// go test -cover ./...
func TestUserEntity(t *testing.T) {
	ctrl := gomock.NewController(t) // 初始化 controller
	defer ctrl.Finish()

	userRepo := NewMockUserRepoInterface(ctrl) // 初始化 mock

	userInst := userEntity{
		userRepo: userRepo,
	}

	Convey("Convey Test User Entity", t, func() {
		var req *vo.UserCreateReq
		var id int
		var err error

		var gender *int
		gender = new(int)
		*gender = 1
		req = &vo.UserCreateReq{
			Name:   "",
			Gender: gender,
		}

		Convey("AddUser Success", func() {
			id = 1
			err = nil
			userRepo.EXPECT().Create(gomock.Any()).Return(id, err)
			idRes, errRes := userInst.AddUser(req)
			assert.Equal(t, idRes, id)
			assert.Equal(t, errRes, err)
		})
		Convey("AddUser Failure", func() {
			id = 0
			err = errors.New("")
			userRepo.EXPECT().Create(gomock.Any()).Return(id, err)
			idRes, errRes := userInst.AddUser(req)
			assert.Equal(t, idRes, id)
			assert.Equal(t, errRes, err)
		})
	})
}
