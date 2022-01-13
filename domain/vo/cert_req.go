package vo

import "time"

// CertCreateReq .
type CertCreateReq struct {
	AuthId             string    `binding:"required" form:"auth_id" json:"auth_id"`                         // 客户端ID
	PVersion           string    `binding:"required" form:"p_version" json:"p_version"`                     // 接口版本
	ContRep            string    `binding:"required" form:"cont_rep" json:"cont_rep"`                       // 内容存储库
	SerialNumber       string    `binding:"required" form:"serial_number" json:"serial_number"`             // 证书序列号
	Version            *int      `binding:"required" form:"version" json:"version"`                         // 证书版本（0:V1,1:V2,2:V3）
	IssuerName         string    `binding:"required" form:"issuer_name" json:"issuer_name"`                 // 颁发机构
	SignatureAlgorithm string    `binding:"required" form:"signature_algorithm" json:"signature_algorithm"` // 签名算法
	NotBefore          time.Time `binding:"required" form:"not_before" json:"not_before"`                   // 有效期开始时间
	NotAfter           time.Time `binding:"required" form:"not_after" json:"not_after"`                     // 有效期结束时间
	EnabledState       *int      `binding:"required" form:"enabled_state" json:"enabled_state"`             // 启用状态（0:已停用,1:已启用）
}

// CertGetListReq .
type CertGetListReq struct {
	Limit    int    `binding:"omitempty" form:"limit,omitempty" json:"limit,omitempty"`
	Offset   int    `binding:"omitempty" form:"offset,omitempty" json:"offset,omitempty"`
	AuthId   string `binding:"omitempty" form:"auth_id,omitempty" json:"auth_id,omitempty"`     // 客户端ID
	PVersion string `binding:"omitempty" form:"p_version,omitempty" json:"p_version,omitempty"` // 接口版本
	ContRep  string `binding:"omitempty" form:"cont_rep,omitempty" json:"cont_rep,omitempty"`   // 内容仓库
}

// CertUpdateReq .
type CertUpdateReq struct {
	Name         string `binding:"omitempty" json:"name"`
	Gender       int    `binding:"omitempty" json:"gender"`
	EnabledState string `binding:"omitempty" json:"enabled_state"`
}
