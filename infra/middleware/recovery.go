package middleware

import (
	"github.com/gin-gonic/gin"
	"sap_cert_mgt/infra/errors"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 错误处理
		defer func() {
			for _, err := range c.Errors {
				statusCode := c.Writer.Status()

				switch interface{}(err.Err).(type) {
				// 接口错误
				case *errors.ApiError:
					c.AbortWithStatusJSON(statusCode, err.Err)
				// 组件错误 todo
				// 系统异常
				default:
					c.AbortWithStatusJSON(statusCode, gin.H{
						"err_code": errors.ErrCode,
						"err_msg": err.Error(),
					})
				}
				return
			}
		}()

		//contentType := c.ContentType()
		//if contentType != "" {
		//	c.Writer.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", c.ContentType()))
		//}
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}
