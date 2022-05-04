package vo

type NoticeConfGetEmailRes struct {
	ServerHost string `json:"server_host"`
	ServerPort int    `json:"server_port"`
	FromName   string `json:"from_name"`
	FromEmail  string `json:"from_email"`
	FromPasswd string `json:"from_passwd"`
}
