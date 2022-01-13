package vo

// NoticeConfModEmailReq .
type NoticeConfModEmailReq struct {
	ServerHost string `binding:"required" json:"server_host"`
	ServerPort string `binding:"required" json:"server_port"`
	FromName   string `binding:"required" json:"from_name"`
	FromEmail  string `binding:"required" json:"from_email"`
	FromPasswd string `binding:"required" json:"from_passwd"`
}
