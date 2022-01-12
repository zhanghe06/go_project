package enums

type NoticeType int

// 通知类型（0:邮件,1:短信）
const (
	NoticeTypeEmail = iota
	NoticeTypeSMS
)

var NoticeTypeMap = map[NoticeType]string{
	NoticeTypeEmail: "邮件",
	NoticeTypeSMS:   "短信",
}

func (t NoticeType) DisplayName() string {
	return NoticeTypeMap[t]
}
