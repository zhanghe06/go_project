package errors

import (
	"encoding/json"
)

var (
	Languages = [3]string{"zh_CN", "zh_TW", "en_US"}
	Lang = "en_US"

)

//func SetLanguage(lang string)  {
//	for
//}

type ApiError struct {
	ErrCode int `json:"code"`
	ErrMsg string `json:"description"`
	ErrSolution string `json:"solution"`
	ErrorLink string `json:"link"`
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
	ErrCodeUnauthorized
	ErrCodeInvalidRequest = 400000000
	ErrCodeNotAuthorized = 401000000
	ErrCodeNotPermission = 403000000
	ErrCodeNotFound = 404000000
	ErrCodeConflicts = 409000000
	ErrCodeInternalServerError = 500000000
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
	ErrMsgInternalServerError = "InternalServerError"
	ErrMsgUnauthorized = "Unauthorized"
)

const (
	// ErrMsgUser 用户模块
	ErrMsgUser = ""
	ErrMsgUserNotFound = "UserNotFound"
	ErrMsgUserDisabled = "UserDisabled"
)

const (
	// ErrMsgCert 证书模块
	ErrMsgCert = ""
	ErrMsgCertNotFound = "CertNotFound"
	ErrMsgCertDisabled = "CertDisabled"
)

const (
	// ErrMsgNoticeConf 通知配置
	ErrMsgNoticeConf = ""
	ErrMsgNoticeConfNotFound = "NoticeConfNotFound"
	ErrMsgNoticeConfDisabled = "NoticeConfDisabled"
)

const (
	// ErrMsgNoticeStrategy 通知策略
	ErrMsgNoticeStrategy = ""
	ErrMsgNoticeStrategyNotFound = "NoticeStrategyNotFound"
	ErrMsgNoticeStrategyDisabled = "NoticeStrategyDisabled"
)

const (
	// ErrMsgOperationLog 操作日志
	ErrMsgOperationLog = ""
	ErrMsgOperationLogNotFound = "OperationLogNotFound"
	ErrMsgOperationLogDisabled = "OperationLogDisabled"
)

//var (
//	errorI18n = map[int]map[string]string {
//
//	}
//)