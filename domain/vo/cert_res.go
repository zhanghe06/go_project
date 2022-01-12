package vo

import (
	"go_project/domain/enums"
)

// CertCreateRes .
type CertCreateRes struct {
	Id int // 用户id
}

// CertGetInfoRes .
type CertGetInfoRes struct {
	Id                int    `json:"id"`                  // 用户id
	Name              string `json:"name"`                // 名称
	Gender            int    `json:"gender"`              // 性别
	EnabledState      int    `json:"enabled_state"`       // 启用状态
	GenderDisplayName string `json:"gender_display_name"` // 性别（显示名称）
}

func (res *CertGetInfoRes) SetGenderDisplayName() {
	res.GenderDisplayName = enums.GenderType(res.Gender).DisplayName()
}

// CertGetListRes .
type CertGetListRes struct {
	Data  []*CertGetInfoRes
	Count int64
}
