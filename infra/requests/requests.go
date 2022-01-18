package requests

import (
	"github.com/gin-gonic/gin"
	"sap_cert_mgt/infra/errors"
	"strings"
)

type IDUriReq struct {
	ID int `uri:"id" binding:"required"`
}

type UserInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TokenAuthorization(c *gin.Context) (userInfo *UserInfo, err error) {
	tokenID := c.GetHeader("Authorization")
	token := strings.TrimPrefix(tokenID, "Bearer ")
	if token == "" {
		err = &errors.ApiError{
			ErrCode: errors.ErrCodeNotAuthorized,
			ErrMsg:  errors.ErrMsgUnauthorized,
		}
		//_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	// todo check token
	userInfo = &UserInfo{
		"1",
		"admin",
	}
	return
}
