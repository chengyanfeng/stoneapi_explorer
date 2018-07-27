package models

type Node struct {
	Data    string `json:"data"`
	TxIds   string `bson:"txIds"`
	AppKey  string `bson:"appKey"`
	Created string `bson:"created"`
	__v     string `bson:"__v";`
	Number  int    `json:"number"`
}
