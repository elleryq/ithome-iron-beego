package controllers

import (
	"github.com/astaxie/beego"
)

// LogoutController operations for Logout
type LogoutController struct {
	beego.Controller
}

// Get
func (c *LogoutController) Get() {
	c.DelSession("user_id")
	c.TplName = "logout/index.tpl"
}
