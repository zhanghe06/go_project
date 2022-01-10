package vo

// UserCreateReq .
type UserCreateReq struct {
	Name   string `binding:"required" form:"name" json:"name"`     // 姓名
	Gender *int   `binding:"required" form:"gender" json:"gender"` // 性别(允许零值，需要设置指针类型)
}

// UserGetListReq .
type UserGetListReq struct {
	Limit  int    `binding:"omitempty" form:"limit,omitempty" json:"limit,omitempty"`
	Offset int    `binding:"omitempty" form:"offset,omitempty" json:"offset,omitempty"`
	Name   string `binding:"omitempty" form:"name,omitempty" json:"name,omitempty"`
}

// UserUpdateReq .
type UserUpdateReq struct {
	Name         string `binding:"omitempty" json:"name"`
	Gender       int    `binding:"omitempty" json:"gender"`
	EnabledState string `binding:"omitempty" json:"enabled_state"`
}
