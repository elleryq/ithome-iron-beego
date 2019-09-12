package controllers

import (
	"github.com/astaxie/beego"
)

// AboutController operations for About
type AboutController struct {
	beego.Controller
}

// URLMapping ...
func (c *AboutController) URLMapping() {
	c.Mapping("Get", c.Get)
}

func (c *AboutController) Get() {
	c.Data["Name"] = "John Doe"
	c.TplName = "about.tpl"
}
