package vo

// NoticeStrategyCreateReq .
type NoticeStrategyCreateReq struct {
	NoticeType       *int   `binding:"required" form:"notice_type" json:"notice_type"`             // 通知类型（0:邮件,1:短信）(允许零值，需要设置指针类型)
	TriggerThreshold *int   `binding:"required" form:"trigger_threshold" json:"trigger_threshold"` // 触发阈值 (允许零值，需要设置指针类型)
	EnabledState     *int   `binding:"required" form:"enabled_state" json:"enabled_state"`         // 启用状态（0:停用,1:启用）
	ToEmails         string `binding:"required" form:"to_emails" json:"to_emails"`                 // 接收邮箱 (半角逗号分隔)
	//CreatedBy        string `binding:"omitempty" form:"created_by,omitempty" json:"created_by,omitempty"` // 创建人员
}

// NoticeStrategyGetListReq .
type NoticeStrategyGetListReq struct {
	Limit        *int     `binding:"omitempty" form:"limit,omitempty" json:"limit,omitempty"`
	Offset       *int     `binding:"omitempty" form:"offset,omitempty" json:"offset,omitempty"`
	NoticeType   *int     `binding:"omitempty" form:"notice_type,omitempty" json:"notice_type,omitempty"`     // 通知类型（0:邮件,1:短信）(允许零值，需要设置指针类型)
	EnabledState *int     `binding:"omitempty" form:"enabled_state,omitempty" json:"enabled_state,omitempty"` // 启用状态（0:停用,1:启用）
	CreatedAt    []string `binding:"omitempty" form:"created_at[],omitempty" json:"created_at[],omitempty"`   // 创建时间
	Sorter       string   `binding:"omitempty" form:"sorter,omitempty" json:"sorter,omitempty"`               // 排序
}

// NoticeStrategyUpdateReq .
type NoticeStrategyUpdateReq struct {
	NoticeType       *int   `binding:"omitempty" json:"notice_type,omitempty"`       // 通知类型（0:邮件,1:短信）(允许零值，需要设置指针类型)
	TriggerThreshold *int   `binding:"omitempty" json:"trigger_threshold,omitempty"` // 触发阈值 (允许零值，需要设置指针类型)
	EnabledState     *int   `binding:"omitempty" json:"enabled_state,omitempty"`     // 启用状态（0:停用,1:启用）
	ToEmails         string `binding:"omitempty" json:"to_emails,omitempty"`         // 接收邮箱 (半角逗号分隔)
	//UpdatedBy        string `binding:"omitempty" form:"updated_by,omitempty" json:"updated_by,omitempty"` // 更新人员
}
