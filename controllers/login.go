package controllers

import (
	"html/template"
	"my/hello/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// Post ...
func (c *LoginController) Post() {
	var msg string
	flash := beego.NewFlash()

	c.TplName = "login/index.tpl"
	// Check XSRF first.
	if !c.CheckXSRFCookie() {
		msg = "XSRF token missing or incorrect."
		c.Data["error"] = msg
		flash.Error(msg)
		flash.Store(&c.Controller)
		return
	}

	// parse form parameters
	username := c.GetString("username")
	password := c.GetString("password")

	o := orm.NewOrm()
	v := &models.Member{Username: username}
	if err := o.QueryTable(new(models.Member)).Filter("Username", username).RelatedSel().One(v); err != nil {
		// not found
		msg = "No such member."
		c.Data["error"] = msg
		flash.Error(msg)
		flash.Store(&c.Controller)
		return
	}
	// found, check password
	if password == v.Password {
		c.SetSession("user_id", int(v.Id))
	} else {
		msg = "Authentication fail."
		c.Data["error"] = msg
		flash.Error(msg)
		flash.Store(&c.Controller)
		return
	}
	flash.Success("Login successed.")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, "/myuser/")
}

// Get ...
func (c *LoginController) Get() {
	flash := beego.NewFlash()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login/index.tpl"
	flash.Store(&c.Controller)
}
