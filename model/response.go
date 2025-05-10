package model

type Resp struct {
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
