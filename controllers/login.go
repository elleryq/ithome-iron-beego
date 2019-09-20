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
	c.TplName = "login/index.tpl"
	c.Data["has_error"] = false
	// Check XSRF first.
	if !c.CheckXSRFCookie() {
		c.Data["error"] = "XSRF token missing or incorrect."
		c.Data["has_error"] = true
		return
	}

	// parse form parameters
	username := c.GetString("username")
	password := c.GetString("password")

	o := orm.NewOrm()
	v := &models.Member{Username: username}
	if err := o.QueryTable(new(models.Member)).Filter("Username", username).RelatedSel().One(v); err != nil {
		// not found
		c.Data["error"] = "No such member."
		c.Data["has_error"] = true
		return
	}
	// found, check password
	if password == v.Password {
		c.SetSession("user_id", int(v.Id))
	} else {
		c.Data["error"] = "Authentication fail."
		c.Data["has_error"] = true
		return
	}
	c.Ctx.Redirect(302, "/myuser/")
}

// Get ...
func (c *LoginController) Get() {
	c.Data["has_error"] = false
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login/index.tpl"
}
