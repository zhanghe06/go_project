package model

import (
	"time"
)

// NoticeStrategy 通知策略
type NoticeStrategy struct {
	Id               int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	NoticeType       int       `gorm:"column:notice_type;type:tinyint(4);default:0;NOT NULL" json:"notice_type"`                // 证书过期通知类型（0:邮件,1:短信）
	TriggerThreshold int       `gorm:"column:trigger_threshold;type:tinyint(4);default:0;NOT NULL" json:"trigger_threshold"`    // 证书过期触发阈值（单位:天;范围:0-255）
	ToEmails         string    `gorm:"column:to_emails;type:varchar(256);NOT NULL" json:"to_emails"`                            // 接收邮箱（半角逗号分隔）
	EnabledState     int       `gorm:"column:enabled_state;type:tinyint(4);default:0;NOT NULL" json:"enabled_state"`            // 启用状态（0:已停用,1:已启用）
	DeletedState     int       `gorm:"column:deleted_state;type:tinyint(4);default:0;NOT NULL" json:"deleted_state"`            // 删除状态（0:未删除,1:已删除）
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"updated_at"` // 更新时间
	DeletedAt        time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                                      // 删除时间
	CreatedBy        string    `gorm:"column:created_by;type:varchar(64);NOT NULL" json:"created_by"`                           // 创建人员
	UpdatedBy        string    `gorm:"column:updated_by;type:varchar(64);NOT NULL" json:"updated_by"`                           // 更新人员
	DeletedBy        string    `gorm:"column:deleted_by;type:varchar(64);NOT NULL" json:"deleted_by"`                           // 删除人员
}

func (m *NoticeStrategy) TableName() string {
	return "notice_strategy"
}
