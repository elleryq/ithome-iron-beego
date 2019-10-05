package controllers

import (
	"html/template"

	"github.com/elleryq/ithome-iron-beego/models"

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
	var username, password string
	var v *models.Member
	var err error

	flash := beego.NewFlash()
	o := orm.NewOrm()

	c.TplName = "login/index.tpl"
	// Check XSRF first.
	if !c.CheckXSRFCookie() {
		msg = "XSRF token missing or incorrect."
		goto ERROR
	}

	// parse form parameters
	username = c.GetString("username")
	password = c.GetString("password")

	v = &models.Member{Username: username}
	if err = o.QueryTable(new(models.Member)).Filter("Username", username).RelatedSel().One(v); err != nil {
		// not found
		msg = "No such member."
		goto ERROR
	}
	// found, check password
	if password == v.Password {
		c.SetSession("user_id", int(v.Id))
	} else {
		msg = "Authentication fail."
		goto ERROR
	}
	flash.Success("Login successed.")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, beego.URLFor("PostController.GetAll"))
	goto FINAL

ERROR:
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["error"] = msg
	flash.Error(msg)
	flash.Store(&c.Controller)

FINAL:
}

// Get ...
func (c *LoginController) Get() {
	flash := beego.NewFlash()
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login/index.tpl"
	flash.Store(&c.Controller)
}
