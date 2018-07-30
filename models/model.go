package models

import "time"

type Node struct {
	Data    string `json:"data"`
	TxIds   string `bson:"txIds"`
	AppKey  string `bson:"appKey"`
	Created time.Time `bson:"created"`
	__v     string `bson:"__v";`
	Number  int    `json:"number"`
}
