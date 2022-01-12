package enums

type EventState int

// 事件状态（0:pending准备,1:waiting等待,2:process进行,3:success成功,4:failure失败）
const (
	Pending = iota
	Waiting
	Process
	Success
	Failure
)

var EventStateMap = map[EventState]string{
	Pending: "准备",
	Waiting: "等待",
	Process: "进行",
	Success: "成功",
	Failure: "失败",
}

func (t EventState) DisplayName() string {
	return EventStateMap[t]
}
