package event

type NoticeInterface interface {
	Scan()
	Send()
	//Send(receivers, subject, body string) error
}
