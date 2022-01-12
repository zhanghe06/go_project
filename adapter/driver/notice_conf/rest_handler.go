package noticeConf

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_project/adapter/driver"
	"go_project/domain/entity"
	"go_project/domain/vo"
	"go_project/infra/logs"
	"go_project/infra/responses"
	"net/http"
	"sync"
)

type restHandler struct {
	noticeConfEntity entity.NoticeConfEntityInterface
	log              logs.Logger
}

var (
	restOnce sync.Once
	restHand driver.RESTHandlerInterface
)

func NewRESTHandler() driver.RESTHandlerInterface {
	restOnce.Do(func() {
		restHand = &restHandler{
			noticeConfEntity: entity.NewNoticeConfEntity(),
			log:              logs.NewLogger(),
		}
	})
	return restHand
}

func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/notice_conf/email", h.getEmailHandler)
	engine.PUT("/notice_conf/email", h.modEmailHandler)
}

func (h *restHandler) getEmailHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 逻辑处理
	data, err := h.noticeConfEntity.GetNoticeConfEmail()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, data)
}

func (h *restHandler) modEmailHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 逻辑处理
	var noticeConfUpdateReq vo.NoticeConfModEmailReq
	if err := c.ShouldBindJSON(&noticeConfUpdateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 参数转换
	var data map[string]interface{}
	reqBytes, err := json.Marshal(noticeConfUpdateReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(reqBytes, &data)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.noticeConfEntity.ModNoticeConfEmail(data)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 响应处理
	c.Status(http.StatusNoContent)
}
