package cert

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"sap_cert_mgt/adapter/driver"
	"sap_cert_mgt/domain/entity"
	"sap_cert_mgt/domain/vo"
	"sap_cert_mgt/infra/errors"
	"sap_cert_mgt/infra/logs"
	"sap_cert_mgt/infra/requests"
	"sap_cert_mgt/infra/responses"
	"net/http"
	"sync"
)

type restHandler struct {
	certEntity entity.CertEntityInterface
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
			certEntity: entity.NewCertEntity(),
			operationLogEntity: entity.NewOperationLogEntity(),
			log:  logs.NewLogger(),
		}
	})
	return restHand
}

func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/cert", h.getListHandler)
	engine.POST("/cert", h.createHandler)
	engine.GET("/cert/:id", h.getInfoHandler)
	engine.DELETE("/cert/:id", h.deleteHandler)
}


func (h *restHandler) getListHandler(c *gin.Context) {
	// 异常捕获
	defer responses.ApiRecover(c)

	// 请求处理
	var certGetListReq vo.CertGetListReq
	if err := c.ShouldBindQuery(&certGetListReq); err != nil {
		apiErr := &errors.ApiError{
			ErrCode: errors.ErrCodeCert,
			ErrMsg:  errors.ErrMsgCert,
		}
		_ = c.AbortWithError(http.StatusBadRequest, apiErr)
		return
	}

	// 参数转换
	var filter map[string]interface{}
	certGetListReqBytes, err := json.Marshal(certGetListReq)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(certGetListReqBytes, &filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	total, data, err := h.certEntity.GetCertList(filter)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.JSON(http.StatusOK, gin.H{
		"total_count": total,
		"entries":  data,
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
	data, err := h.certEntity.GetCertInfo(uriReq.ID)
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

	// 请求处理
	var certCreateReq vo.CertCreateReq
	if err := c.ShouldBindJSON(&certCreateReq); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 逻辑处理
	id, err := h.certEntity.AddCert(&certCreateReq, "SAP")
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Header("Location", c.FullPath() + fmt.Sprintf("/%d", id))
	c.JSON(http.StatusCreated, gin.H{
		"id":  id,
	})
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
	err = h.certEntity.DelCert(uriReq.ID, userInfo.ID)
	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// 响应处理
	c.Status(http.StatusNoContent)
}
