package routers

import (
	"stoneapi_explorer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/getdata", &controllers.MainController{},"get:Getdata")

}
