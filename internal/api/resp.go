package api

type Entity struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Msg       string      `json:"msg"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
	Data      interface{} `json:"data"`
}
