package driver

import (
	"github.com/gin-gonic/gin"
)

// RESTHandlerInterface .
type RESTHandlerInterface interface {
	// RegisterAPI 注册API
	RegisterAPI(engine *gin.Engine)
}
