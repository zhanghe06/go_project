package enums

type EnabledState int

// 启用状态（0:停用,1:启用）
const (
	Disabled = iota
	Enabled
)

var enabledStateMap = map[EnabledState]string{
	Disabled: "停用",
	Enabled:  "启用",
}

func (t EnabledState) DisplayName() string {
	return enabledStateMap[t]
}
