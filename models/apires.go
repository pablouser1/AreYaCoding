package models

type ApiRes struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
