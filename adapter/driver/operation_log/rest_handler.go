package operationLog

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
	operationLogEntity entity.OperationLogEntityInterface
	log logs.Logger
}

var (
	restOnce sync.Once
	restHand driver.RESTHandlerInterface
)

func NewRESTHandler() driver.RESTHandlerInterface {
	restOnce.Do(func() {
		restHand = &restHandler{
			operationLogEntity: entity.NewOperationLogEntity(),
			log:  logs.NewLogger(),
		}
	})
	return restHand
}

func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/operation_log", h.getListHandler)
	engine.POST("/operation_log", h.createHandler)
	engine.GET("/operation_log/:id", h.getInfoHandler)
	engine.PUT("/operation_log/:id", h.updateHandler)
	engine.DELETE("/operation_log/:id", h.deleteHandler)
}


func (h *restHandler) getListHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 请求处理
	var operationLogGetListReq vo.OperationLogGetListReq
	if err := c.ShouldBindQuery(&operationLogGetListReq); err != nil {
		apiErr := &errors.ApiError{
			ErrCode: errors.ErrCodeOperationLog,
			ErrMsg:  errors.ErrMsgOperationLog,
		}
		_ = c.AbortWithError(http.StatusBadRequest, apiErr)
		return
	}

	// 参数转换
	var filter map[string]interface{}
	operationLogGetListReqBytes, err := json.Marshal(operationLogGetListReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(operationLogGetListReqBytes, &filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	total, data, err := h.operationLogEntity.GetOperationLogList(filter)
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
	data, err := h.operationLogEntity.GetOperationLogInfo(uriReq.ID)
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
	var operationLogCreateReq vo.OperationLogCreateReq
	if err := c.ShouldBindJSON(&operationLogCreateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	id, err := h.operationLogEntity.AddOperationLog(&operationLogCreateReq)
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

	var operationLogUpdateReq vo.OperationLogUpdateReq
	if err := c.ShouldBindJSON(&operationLogUpdateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 参数转换
	var data map[string]interface{}
	reqBytes, err := json.Marshal(operationLogUpdateReq)
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
	err = h.operationLogEntity.ModOperationLog(uriReq.ID, data)
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
	err := h.operationLogEntity.DelOperationLog(uriReq.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Status(204)
}
