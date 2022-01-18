package model

import (
	"time"
)

// OperationLog 操作日志
type OperationLog struct {
	Id           int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	OpType       string    `gorm:"column:op_type;type:varchar(20);NOT NULL" json:"op_type"`                                 // 操作类型（create:创建、update:更新、delete:删除）
	RsType       string    `gorm:"column:rs_type;type:varchar(20);NOT NULL" json:"rs_type"`                                 // 资源类型（cert、）
	RsId         int       `gorm:"column:rs_id;type:int(11);default:0;NOT NULL" json:"rs_id"`                               // 资源ID
	OpDetail     string    `gorm:"column:op_detail;type:varchar(512);NOT NULL" json:"op_detail"`                            // 操作详情
	OpError      string    `gorm:"column:op_error;type:varchar(512);NOT NULL" json:"op_error"`                              // 操作错误
	DeletedState int       `gorm:"column:deleted_state;type:tinyint(4);default:0;NOT NULL" json:"deleted_state"`            // 删除状态（0:未删除,1:已删除）
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"updated_at"` // 更新时间
	DeletedAt    time.Time `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                                      // 删除时间
	CreatedBy    string    `gorm:"column:created_by;type:varchar(64);NOT NULL" json:"created_by"`                           // 创建人员
	UpdatedBy    string    `gorm:"column:updated_by;type:varchar(64);NOT NULL" json:"updated_by"`                           // 更新人员
	DeletedBy    string    `gorm:"column:deleted_by;type:varchar(64);NOT NULL" json:"deleted_by"`                           // 删除人员
}

func (m *OperationLog) TableName() string {
	return "operation_log"
}
