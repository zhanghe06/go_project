package model

import (
	"time"
)

// NoticeEvent 通知事件
type NoticeEvent struct {
	Id               int          `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	CertId           int          `gorm:"column:cert_id;type:int(11);default:0;NOT NULL" json:"cert_id"`                           // 证书ID
	NoticeStrategyId int          `gorm:"column:notice_strategy_id;type:int(11);default:0;NOT NULL" json:"notice_strategy_id"`     // 策略ID
	EventState       int          `gorm:"column:event_state;type:tinyint(4);default:0;NOT NULL" json:"event_state"`                // 事件状态（0:pending准备,1:waiting等待,2:process进行,3:success成功,4:failure失败）
	DeletedState     int          `gorm:"column:deleted_state;type:tinyint(4);default:0;NOT NULL" json:"deleted_state"`            // 删除状态（0:未删除,1:已删除）
	CreatedAt        time.Time    `gorm:"column:created_at;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt        time.Time    `gorm:"column:updated_at;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"updated_at"` // 更新时间
	DeletedAt        time.Time     `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                                     // 删除时间
	CreatedBy        string       `gorm:"column:created_by;type:varchar(64);NOT NULL" json:"created_by"`                           // 创建人员
	UpdatedBy        string       `gorm:"column:updated_by;type:varchar(64);NOT NULL" json:"updated_by"`                           // 更新人员
	DeletedBy        string       `gorm:"column:deleted_by;type:varchar(64);NOT NULL" json:"deleted_by"`                           // 删除人员
}

func (m *NoticeEvent) TableName() string {
	return "notice_event"
}
