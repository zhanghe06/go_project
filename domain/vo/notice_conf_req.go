package vo

// NoticeConfModEmailReq .
type NoticeConfModEmailReq struct {
	ServerHost string `binding:"omitempty" form:"server_host,omitempty" json:"server_host,omitempty"`
	ServerPort string `binding:"omitempty" form:"server_port,omitempty" json:"server_port,omitempty"`
	FromName   string `binding:"omitempty" form:"from_name,omitempty" json:"from_name,omitempty"`
	FromEmail  string `binding:"omitempty" form:"from_email,omitempty" json:"from_email,omitempty"`
	FromPasswd string `binding:"omitempty" form:"from_passwd,omitempty" json:"from_passwd,omitempty"`
}
