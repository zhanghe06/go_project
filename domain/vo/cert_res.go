package vo

import (
	"go_project/domain/enums"
	"time"
)

// CertCreateRes .
type CertCreateRes struct {
	Id int // 证书id
}

// CertGetInfoRes .
type CertGetInfoRes struct {
	Id                      int       `json:"id"`                         // 证书id
	AuthId                  string    `json:"auth_id"`                    // 客户端ID
	PVersion                string    `json:"p_version"`                  // 接口版本
	ContRep                 string    `json:"cont_rep"`                   // 内容存储库
	SerialNumber            string    `json:"serial_number"`              // 证书序列号
	Version                 int       `json:"version"`                    // 证书版本（0:V1,1:V2,2:V3）
	IssuerName              string    `json:"issuer_name"`                // 颁发机构
	SignatureAlgorithm      string    `json:"signature_algorithm"`        // 签名算法
	NotBefore               time.Time `json:"not_before"`                 // 有效期开始时间
	NotAfter                time.Time `json:"not_after"`                  // 有效期结束时间
	EnabledState            int       `json:"enabled_state"`              // 启用状态（0:已停用,1:已启用）
	VersionDisplayName      string    `json:"version_display_name"`       // 证书版本（显示名称）
	EnabledStateDisplayName string    `json:"enabled_state_display_name"` // 启用状态（显示名称）
}

func (res *CertGetInfoRes) SetVersionDisplayName() {
	res.VersionDisplayName = enums.CertVersion(res.Version).DisplayName()
}

func (res *CertGetInfoRes) SetEnabledStateDisplayName() {
	res.EnabledStateDisplayName = enums.EnabledState(res.EnabledState).DisplayName()
}

// CertGetListRes .
type CertGetListRes struct {
	Data  []*CertGetInfoRes
	Count int64
}
