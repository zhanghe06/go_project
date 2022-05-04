package requests

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"sap_cert_mgt/infra/errors"
	"strings"
)

type IDUriReq struct {
	ID int `uri:"id" binding:"required"`
}

type UserInfo struct {
	ID   string `json:"user_id"`
	Name string `json:"user_name"`
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
	var tokenObject []byte
	tokenObject, err = base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Printf("base64 decode err, %v\n", err)
		err = &errors.ApiError{
			ErrCode: errors.ErrCodeNotAuthorized,
			ErrMsg:  errors.ErrMsgUnauthorized,
		}
		//_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	if err = json.Unmarshal(tokenObject, &userInfo); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		err = &errors.ApiError{
			ErrCode: errors.ErrCodeNotAuthorized,
			ErrMsg:  errors.ErrMsgUnauthorized,
		}
		//_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	//userInfo = &UserInfo{
	//	"admin",
	//	"admin",
	//}
	return
}
