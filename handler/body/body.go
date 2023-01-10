package body

type LoginInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type HomeRes struct {
	UserName string `json:"user_name"`
}
