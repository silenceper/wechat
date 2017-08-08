package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/wechat", &controllers.MainController{})
}
