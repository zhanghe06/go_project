package clients

type NoticeInterface interface {
	Send(receivers, subject, body string) error
}
