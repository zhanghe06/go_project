package user

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
	userEntity entity.UserEntityInterface
	log logs.Logger
}

var (
	restOnce sync.Once
	restHand driver.RESTHandlerInterface
)

func NewRESTHandler() driver.RESTHandlerInterface {
	restOnce.Do(func() {
		restHand = &restHandler{
			userEntity: entity.NewUserEntity(),
			log:  logs.NewLogger(),
		}
	})
	return restHand
}

func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/user", h.getListHandler)
	engine.POST("/user", h.createHandler)
	engine.GET("/user/:id", h.getInfoHandler)
	engine.PUT("/user/:id", h.updateHandler)
	engine.DELETE("/user/:id", h.deleteHandler)
}


func (h *restHandler) getListHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 请求处理
	var userGetListReq vo.UserGetListReq
	if err := c.ShouldBindQuery(&userGetListReq); err != nil {
		apiErr := &errors.ApiError{
			ErrCode: errors.ErrCodeUser,
			ErrMsg:  errors.ErrMsgUser,
		}
		_ = c.AbortWithError(http.StatusBadRequest, apiErr)
		return
	}

	// 参数转换
	var filter map[string]interface{}
	userGetListReqBytes, err := json.Marshal(userGetListReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(userGetListReqBytes, &filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	filterArgs := make([]interface{}, 0)
	// filterArgs = append(filterArgs, "name <> ?")
	// filterArgs = append(filterArgs, "呵呵")
	total, data, err := h.userEntity.GetUserList(filter, filterArgs...)
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
	data, err := h.userEntity.GetUserInfo(uriReq.ID)
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
	var userCreateReq vo.UserCreateReq
	if err := c.ShouldBindJSON(&userCreateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	id, err := h.userEntity.AddUser(&userCreateReq)
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

	var userUpdateReq vo.UserUpdateReq
	if err := c.ShouldBindJSON(&userUpdateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 参数转换
	var data map[string]interface{}
	reqBytes, err := json.Marshal(userUpdateReq)
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
	err = h.userEntity.ModUser(uriReq.ID, data)
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
	err := h.userEntity.DelUser(uriReq.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Status(204)
}
