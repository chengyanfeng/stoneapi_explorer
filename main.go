package main

import (
	_ "redfind/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.HTTPPort = 8080 //端口设置
	beego.Run()
}

