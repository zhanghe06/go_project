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

// 分组自定义错误码，每组使用const分别定义
const (
	// ErrCode 通用异常
	ErrCode = 100000 + iota
	ErrCodeInternalServerError
)

const (
	// ErrCodeUser 用户模块
	ErrCodeUser = 200000 + iota
	ErrCodeUserNotFound
	ErrCodeUserDisabled
)

// 分组自定义错误消息，每组使用const分别定义
const (
	// ErrMsg 通用异常
	ErrMsg = ""
	ErrMsgInternalServerError = ""
)

const (
	// ErrMsgUser 用户模块
	ErrMsgUser = ""
	ErrMsgUserNotFound = ""
	ErrMsgUserDisabled = ""
)
