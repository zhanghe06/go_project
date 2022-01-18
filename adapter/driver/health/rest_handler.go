package health

import (
	"github.com/gin-gonic/gin"
	"sap_cert_mgt/adapter/driver"
	"net/http"
	"sync"
)

var (
	restOnce sync.Once
	restHand driver.RESTHandlerInterface
)

type restHandler struct {
}

// NewRESTHandler 创建公共RESTful api handler对象
func NewRESTHandler() driver.RESTHandlerInterface {
	restOnce.Do(func() {
		restHand = &restHandler{}
	})

	return restHand
}

// RegisterAPI 注册API
func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/health/ready", getHealth)
	engine.GET("/health/alive", getAlive)
}

func getHealth(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, "ready")
}

func getAlive(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, "alive")
}
