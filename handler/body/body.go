package body

import "time"

type LoginInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LeavingMessage struct {
	Title string `json:"title"`
	Msg   string `json:"msg"`
}

type Res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type HomeRes struct {
	UserName string `json:"user_name"`
}

type BuyCarRes struct {
	Image    string  `json:"image"`
	Describe string  `json:"describe"`
	Price    float64 `json:"price"`
	Number   int     `json:"number"`
}

type OrderRes struct {
	ID        int       `json:"id"`
	StartTime time.Time `json:"start_time"`
	Status    int       `json:"status"`
}

type OrderReq struct {
	Pid     int    `json:"pid"`
	Number  int    `json:"num"`
	Address string `json:"address"`
}

type GetProductRea struct {
	Pid int `json:"pod"`
}
