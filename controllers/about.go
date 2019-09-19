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
	key := "asta"
	v := c.GetSession(key)
	if v == nil {
		c.SetSession(key, int(1))
		c.Data[key] = 0
	} else {
		c.SetSession(key, v.(int)+1)
		c.Data[key] = v.(int)
	}

	c.Data["Name"] = "John Doe"
	c.TplName = "about.tpl"
}
