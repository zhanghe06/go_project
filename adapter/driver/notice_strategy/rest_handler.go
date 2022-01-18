package noticeStrategy

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sap_cert_mgt/adapter/driver"
	"sap_cert_mgt/domain/entity"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/errors"
	"sap_cert_mgt/infra/logs"
	"sap_cert_mgt/infra/requests"
	"sap_cert_mgt/infra/responses"
	"sync"
)

type restHandler struct {
	noticeStrategyEntity entity.NoticeStrategyEntityInterface
	log                  logs.Logger
}

var (
	restOnce sync.Once
	restHand driver.RESTHandlerInterface
)

func NewRESTHandler() driver.RESTHandlerInterface {
	restOnce.Do(func() {
		restHand = &restHandler{
			noticeStrategyEntity: entity.NewNoticeStrategyEntity(),
			log:                  logs.NewLogger(),
		}
	})
	return restHand
}

func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/notice_strategy", h.getListHandler)
	engine.POST("/notice_strategy", h.createHandler)
	engine.GET("/notice_strategy/:id", h.getInfoHandler)
	engine.PUT("/notice_strategy/:id", h.updateHandler)
	engine.DELETE("/notice_strategy/:id", h.deleteHandler)
}

func (h *restHandler) getListHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 认证处理
	_, err := requests.TokenAuthorization(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// 请求处理
	var noticeStrategyGetListReq vo.NoticeStrategyGetListReq
	if err := c.ShouldBindQuery(&noticeStrategyGetListReq); err != nil {
		apiErr := &errors.ApiError{
			ErrCode: errors.ErrCodeNoticeStrategy,
			ErrMsg:  errors.ErrMsgNoticeStrategy,
		}
		_ = c.AbortWithError(http.StatusBadRequest, apiErr)
		return
	}

	// 参数转换
	var filter map[string]interface{}
	noticeStrategyGetListReqBytes, err := json.Marshal(noticeStrategyGetListReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(noticeStrategyGetListReqBytes, &filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	total, data, err := h.noticeStrategyEntity.GetNoticeStrategyList(filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, gin.H{
		"total_count": total,
		"entries":     data,
	})
}

func (h *restHandler) getInfoHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 认证处理
	_, err := requests.TokenAuthorization(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// 请求处理
	var uriReq requests.IDUriReq
	if err := c.ShouldBindUri(&uriReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	data, err := h.noticeStrategyEntity.GetNoticeStrategyInfo(uriReq.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, data)
}

func (h *restHandler) createHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 认证处理
	userInfo, err := requests.TokenAuthorization(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// 请求处理
	var noticeStrategyCreateReq vo.NoticeStrategyCreateReq
	if err := c.ShouldBindJSON(&noticeStrategyCreateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	id, err := h.noticeStrategyEntity.AddNoticeStrategy(&noticeStrategyCreateReq, userInfo.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Header("Location", c.FullPath()+fmt.Sprintf("/%d", id))
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *restHandler) updateHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 认证处理
	userInfo, err := requests.TokenAuthorization(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// 请求处理
	var uriReq requests.IDUriReq
	if err := c.ShouldBindUri(&uriReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var noticeStrategyUpdateReq vo.NoticeStrategyUpdateReq
	if err := c.ShouldBindJSON(&noticeStrategyUpdateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 参数转换
	var data map[string]interface{}
	reqBytes, err := json.Marshal(noticeStrategyUpdateReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(reqBytes, &data)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	err = h.noticeStrategyEntity.ModNoticeStrategy(uriReq.ID, data, userInfo.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *restHandler) deleteHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 认证处理
	userInfo, err := requests.TokenAuthorization(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// 请求处理
	var uriReq requests.IDUriReq
	if err := c.ShouldBindUri(&uriReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	err = h.noticeStrategyEntity.DelNoticeStrategy(uriReq.ID, userInfo.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Status(http.StatusNoContent)
}
