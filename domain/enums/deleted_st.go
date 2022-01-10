package enums

type DeletedState int

// 删除状态（0:未删除,1:已删除）
const (
	NotDeleted = iota
	HasDeleted
)

var deletedStateMap = map[DeletedState]string{
	NotDeleted: "未删除",
	HasDeleted: "已删除",
}

func (t DeletedState) DisplayName() string {
	return deletedStateMap[t]
}
