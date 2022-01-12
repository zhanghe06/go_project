package errors

import (
	"encoding/json"
)

type ApiError struct {
	ErrCode int `json:"err_code"`
	ErrMsg string `json:"err_msg"`
}

func (e *ApiError) Error() string {
	es, _ := json.Marshal(e)
	return string(es)
}

var _ error = &ApiError{}

// --------
// 错误码 分组自定义错误码，每组使用const分别定义
// --------
const (
	// ErrCode 通用异常
	ErrCode = 100000 + iota
	ErrCodeInternalServerError
	ErrCodeUnauthorized
)

const (
	// ErrCodeUser 用户模块
	ErrCodeUser = 200000 + iota
	ErrCodeUserNotFound
	ErrCodeUserDisabled
)

const (
	// ErrCodeCert 证书模块
	ErrCodeCert = 300000 + iota
	ErrCodeCertNotFound
	ErrCodeCertDisabled
)

const (
	// ErrCodeNoticeConf 通知配置
	ErrCodeNoticeConf = 400000 + iota
	ErrCodeNoticeConfNotFound
	ErrCodeNoticeConfDisabled
)

const (
	// ErrCodeNoticeStrategy 通知策略
	ErrCodeNoticeStrategy = 500000 + iota
	ErrCodeNoticeStrategyNotFound
	ErrCodeNoticeStrategyDisabled
)

const (
	// ErrCodeOperationLog 操作日志
	ErrCodeOperationLog = 600000 + iota
	ErrCodeOperationLogNotFound
	ErrCodeOperationLogDisabled
)

// --------
// 错误消息 分组自定义错误消息，每组使用const分别定义
// --------
const (
	// ErrMsg 通用异常
	ErrMsg = ""
	ErrMsgInternalServerError = ""
	ErrMsgUnauthorized = ""
)

const (
	// ErrMsgUser 用户模块
	ErrMsgUser = ""
	ErrMsgUserNotFound = ""
	ErrMsgUserDisabled = ""
)

const (
	// ErrMsgCert 证书模块
	ErrMsgCert = ""
	ErrMsgCertNotFound = ""
	ErrMsgCertDisabled = ""
)

const (
	// ErrMsgNoticeConf 通知配置
	ErrMsgNoticeConf = ""
	ErrMsgNoticeConfNotFound = ""
	ErrMsgNoticeConfDisabled = ""
)

const (
	// ErrMsgNoticeStrategy 通知策略
	ErrMsgNoticeStrategy = ""
	ErrMsgNoticeStrategyNotFound = ""
	ErrMsgNoticeStrategyDisabled = ""
)

const (
	// ErrMsgOperationLog 操作日志
	ErrMsgOperationLog = ""
	ErrMsgOperationLogNotFound = ""
	ErrMsgOperationLogDisabled = ""
)
