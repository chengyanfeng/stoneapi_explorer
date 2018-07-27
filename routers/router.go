package routers

import (
	"lian/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/getdata", &controllers.MainController{},"get:Getdata")

}
