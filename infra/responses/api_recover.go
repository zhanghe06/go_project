package responses

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRecover(c *gin.Context) {
	if rec := recover(); rec != nil {
		err := fmt.Errorf("%v", rec)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}
}
