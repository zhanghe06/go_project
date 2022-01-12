package enums

type EventState int

// 事件状态（0:pending准备,1:waiting等待,2:process进行,3:success成功,4:failure失败）
const (
	EventStatePending = iota
	EventStateWaiting
	EventStateProcess
	EventStateSuccess
	EventStateFailure
)

var EventStateMap = map[EventState]string{
	EventStatePending: "准备",
	EventStateWaiting: "等待",
	EventStateProcess: "进行",
	EventStateSuccess: "成功",
	EventStateFailure: "失败",
}

func (t EventState) DisplayName() string {
	return EventStateMap[t]
}
