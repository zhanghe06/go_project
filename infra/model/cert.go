package model

import (
	"time"
)

// Cert 证书管理
type Cert struct {
	Id                 int          `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	AuthId             string       `gorm:"column:auth_id;type:varchar(32);NOT NULL" json:"auth_id"`                                 // 客户端ID
	PVersion           string       `gorm:"column:p_version;type:varchar(32);NOT NULL" json:"p_version"`                             // 接口版本
	ContRep            string       `gorm:"column:cont_rep;type:varchar(32);NOT NULL" json:"cont_rep"`                               // 内容存储库
	SerialNumber       string       `gorm:"column:serial_number;type:varchar(20);NOT NULL" json:"serial_number"`                     // 证书序列号
	Version            int          `gorm:"column:version;type:tinyint(4);default:0;NOT NULL" json:"version"`                        // 证书版本（0:V1,1:V2,2:V3）
	IssuerName         string       `gorm:"column:issuer_name;type:varchar(64);NOT NULL" json:"issuer_name"`                         // 颁发机构
	SignatureAlgorithm string       `gorm:"column:signature_algorithm;type:varchar(32);NOT NULL" json:"signature_algorithm"`         // 签名算法
	NotBefore          time.Time    `gorm:"column:not_before;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"not_before"` // 有效期开始时间
	NotAfter           time.Time    `gorm:"column:not_after;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"not_after"`   // 有效期结束时间
	EnabledState       int          `gorm:"column:enabled_state;type:tinyint(4);default:0;NOT NULL" json:"enabled_state"`            // 启用状态（0:已停用,1:已启用）
	DeletedState       int          `gorm:"column:deleted_state;type:tinyint(4);default:0;NOT NULL" json:"deleted_state"`            // 删除状态（0:未删除,1:已删除）
	CreatedAt          time.Time    `gorm:"column:created_at;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt          time.Time    `gorm:"column:updated_at;type:timestamp;default:0000-00-00 00:00:00;NOT NULL" json:"updated_at"` // 更新时间
	DeletedAt          time.Time     `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`                                     // 删除时间
	CreatedBy          string       `gorm:"column:created_by;type:varchar(64);NOT NULL" json:"created_by"`                           // 创建人员
	UpdatedBy          string       `gorm:"column:updated_by;type:varchar(64);NOT NULL" json:"updated_by"`                           // 更新人员
	DeletedBy          string       `gorm:"column:deleted_by;type:varchar(64);NOT NULL" json:"deleted_by"`                           // 删除人员
}

func (m *Cert) TableName() string {
	return "cert"
}
