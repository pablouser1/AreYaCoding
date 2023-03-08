package models

type ErrorRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
