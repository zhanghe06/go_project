package vo

// OperationLogCreateRes .
type OperationLogCreateRes struct {
	Id int // 日志id
}

// OperationLogGetInfoRes .
type OperationLogGetInfoRes struct {
	Id       int    `json:"id"`        // 日志id
	OpType   string `json:"op_type"`   // 操作类型（create:创建、update:更新、delete:删除）
	RsType   string `json:"rs_type"`   // 资源类型（cert、）
	RsId     int    `json:"rs_id"`     // 资源ID
	OpDetail string `json:"op_detail"` // 操作详情
	OpError  string `json:"op_error"`  // 内容存储库
}

// OperationLogGetListRes .
type OperationLogGetListRes struct {
	Data  []*OperationLogGetInfoRes
	Count int64
}
