package controllers

import (
	"github.com/astaxie/beego"
	"beego/services"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.EnableRender = false
	services.Handle(c.Ctx)
}
