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
	c.Data["has_error"] = false
	c.TplName = "logout/index.tpl"
}
