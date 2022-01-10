package vo

import (
	"go_project/domain/enums"
)

// UserCreateRes .
type UserCreateRes struct {
	Id int // 用户id
}

// UserGetInfoRes .
type UserGetInfoRes struct {
	Id                int    `json:"id"`                  // 用户id
	Name              string `json:"name"`                // 名称
	Gender            int    `json:"gender"`              // 性别
	EnabledState      int    `json:"enabled_state"`       // 启用状态
	GenderDisplayName string `json:"gender_display_name"` // 性别（显示名称）
}

func (res *UserGetInfoRes) SetGenderDisplayName() {
	res.GenderDisplayName = enums.GenderType(res.Gender).DisplayName()
}

// UserGetListRes .
type UserGetListRes struct {
	Data  []*UserGetInfoRes
	Count int64
}
