package vo

// OperationLogCreateReq .
type OperationLogCreateReq struct {
	OpType   string `binding:"required" form:"op_type" json:"op_type"`     // 操作类型（create:创建、update:更新、delete:删除）
	RsType   string `binding:"required" form:"rs_type" json:"rs_type"`     // 资源类型（cert、）
	RsId     int    `binding:"required" form:"rs_id" json:"rs_id"`         // 资源ID
	OpDetail string `binding:"required" form:"op_detail" json:"op_detail"` // 操作详情
	OpError  string `binding:"required" form:"op_error" json:"op_error"`   // 内容存储库
}

// OperationLogGetListReq .
type OperationLogGetListReq struct {
	Limit  int    `binding:"omitempty" form:"limit,omitempty" json:"limit,omitempty"`
	Offset int    `binding:"omitempty" form:"offset,omitempty" json:"offset,omitempty"`
	OpType string `binding:"omitempty" form:"op_type,omitempty" json:"op_type,omitempty"`
	RsType string `binding:"omitempty" form:"rs_type,omitempty" json:"rs_type,omitempty"`
}
