package vo

// UserCreateReq .
type UserCreateReq struct {
	Name   string `binding:"required" form:"name" json:"name"`     // 姓名
	Gender *int   `binding:"required" form:"gender" json:"gender"` // 性别(允许零值，需要设置指针类型)
}

// UserGetListReq .
// http://0.0.0.0:8080/user?limit=1 显示限量数据
// http://0.0.0.0:8080/user?limit=0 显示全量数据
// http://0.0.0.0:8080/user?limit=  显示全量数据
type UserGetListReq struct {
	Limit  *int   `binding:"omitempty" form:"limit,omitempty" json:"limit,omitempty"`
	Offset *int   `binding:"omitempty" form:"offset,omitempty" json:"offset,omitempty"`
	Name   string `binding:"omitempty" form:"name,omitempty" json:"name,omitempty"`
}

// UserUpdateReq .
type UserUpdateReq struct {
	Name         string `binding:"omitempty" json:"name,omitempty"`
	Gender       *int   `binding:"omitempty" json:"gender,omitempty"`
	EnabledState string `binding:"omitempty" json:"enabled_state,omitempty"`
}
