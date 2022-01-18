package noticeConf

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"sap_cert_mgt/adapter/driver"
	"sap_cert_mgt/domain/entity"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/logs"
	"sap_cert_mgt/infra/requests"
	"sap_cert_mgt/infra/responses"
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

	// 认证处理
	_, err := requests.TokenAuthorization(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

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

	// 认证处理
	userInfo, err := requests.TokenAuthorization(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

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

	err = h.noticeConfEntity.ModNoticeConfEmail(data, userInfo.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 响应处理
	c.Status(http.StatusNoContent)
}
