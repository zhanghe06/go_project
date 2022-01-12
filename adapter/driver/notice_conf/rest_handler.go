package noticeConf

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_project/adapter/driver"
	"go_project/domain/entity"
	"go_project/domain/vo"
	"go_project/infra/errors"
	"go_project/infra/logs"
	"go_project/infra/requests"
	"go_project/infra/responses"
	"net/http"
	"sync"
)

type restHandler struct {
	noticeConfEntity entity.NoticeConfEntityInterface
	log logs.Logger
}

var (
	restOnce sync.Once
	restHand driver.RESTHandlerInterface
)

func NewRESTHandler() driver.RESTHandlerInterface {
	restOnce.Do(func() {
		restHand = &restHandler{
			noticeConfEntity: entity.NewNoticeConfEntity(),
			log:  logs.NewLogger(),
		}
	})
	return restHand
}

func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/notice_conf", h.getListHandler)
	engine.POST("/notice_conf", h.createHandler)
	engine.GET("/notice_conf/:id", h.getInfoHandler)
	engine.PUT("/notice_conf/:id", h.updateHandler)
	engine.DELETE("/notice_conf/:id", h.deleteHandler)
}


func (h *restHandler) getListHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 请求处理
	var noticeConfGetListReq vo.NoticeConfGetListReq
	if err := c.ShouldBindQuery(&noticeConfGetListReq); err != nil {
		apiErr := &errors.ApiError{
			ErrCode: errors.ErrCodeNoticeConf,
			ErrMsg:  errors.ErrMsgNoticeConf,
		}
		_ = c.AbortWithError(http.StatusBadRequest, apiErr)
		return
	}

	// 参数转换
	var filter map[string]interface{}
	noticeConfGetListReqBytes, err := json.Marshal(noticeConfGetListReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(noticeConfGetListReqBytes, &filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	total, data, err := h.noticeConfEntity.GetNoticeConfList(filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"data":  data,
	})
}

func (h *restHandler) getInfoHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 请求处理
	var uriReq requests.IDUriReq
	if err := c.ShouldBindUri(&uriReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	data, err := h.noticeConfEntity.GetNoticeConfInfo(uriReq.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, gin.H{
		"data":  data,
	})
}

func (h *restHandler) createHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 请求处理
	var noticeConfCreateReq vo.NoticeConfCreateReq
	if err := c.ShouldBindJSON(&noticeConfCreateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	id, err := h.noticeConfEntity.AddNoticeConf(&noticeConfCreateReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusCreated, gin.H{
		"id":  id,
	})
}


func (h *restHandler) updateHandler(c *gin.Context) {
	defer responses.ApiRecover(c)

	// 请求处理
	var uriReq requests.IDUriReq
	if err := c.ShouldBindUri(&uriReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var noticeConfUpdateReq vo.NoticeConfUpdateReq
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

	// 逻辑处理
	err = h.noticeConfEntity.ModNoticeConf(uriReq.ID, data)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.Status(204)
}


func (h *restHandler) deleteHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 请求处理
	var uriReq requests.IDUriReq
	if err := c.ShouldBindUri(&uriReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	err := h.noticeConfEntity.DelNoticeConf(uriReq.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Status(204)
}
