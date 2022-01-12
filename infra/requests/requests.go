package requests

import (
	"github.com/gin-gonic/gin"
	"go_project/infra/errors"
	"strings"
)

type IDUriReq struct {
	ID int `uri:"id" binding:"required"`
}

func TokenAuthorization(c *gin.Context) (err error) {
	tokenID := c.GetHeader("Authorization")
	token := strings.TrimPrefix(tokenID, "Bearer ")
	if token == "" {
		err = &errors.ApiError{
			ErrCode: errors.ErrCodeUnauthorized,
			ErrMsg:  errors.ErrMsgUnauthorized,
		}
		//_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	// todo check token
	return
}
