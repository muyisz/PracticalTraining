package body

type LoginInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginRes struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}
