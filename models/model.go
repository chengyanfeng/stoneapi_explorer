package models

type Node struct {
	Data    string `json:"data";"size:512;column:data"`
	TxIds   string `json:"txids";"size:512;column:txids"`
	AppKey  string `json:"appkey";"size:512;column:appkey"`
	Created string `json:"created";"size:512;column:created"`
	__v     string `json:"__v";"size:512;column:__v"`
}
