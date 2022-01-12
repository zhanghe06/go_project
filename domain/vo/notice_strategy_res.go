package vo

import (
	"go_project/domain/enums"
)

// NoticeStrategyCreateRes .
type NoticeStrategyCreateRes struct {
	Id int // id
}

// NoticeStrategyGetInfoRes .
type NoticeStrategyGetInfoRes struct {
	Id                      int    `json:"id"`                         // id
	NoticeType              int    `json:"notice_type"`                // 通知类型（0:邮件,1:短信）
	TriggerThreshold        int    `json:"trigger_threshold"`          // 触发阈值
	EnabledState            int    `json:"enabled_state"`              // 启用状态（0:停用,1:启用）
	ToEmails                string `json:"to_emails"`                  // 接收邮箱 (半角逗号分隔)
	EnabledStateDisplayName string `json:"enabled_state_display_name"` // 接收邮箱 (半角逗号分隔)
}

func (res *NoticeStrategyGetInfoRes) SetEnabledStateDisplayName() {
	res.EnabledStateDisplayName = enums.EnabledState(res.EnabledState).DisplayName()
}

// NoticeStrategyGetListRes .
type NoticeStrategyGetListRes struct {
	Data  []*NoticeStrategyGetInfoRes
	Count int64
}
