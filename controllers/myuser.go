package controllers

import (
	"html/template"
	"strings"
	"time"

	"github.com/elleryq/ithome-iron-beego/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//  UserController operations for User
type MyUserController struct {
	beego.Controller
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *MyUserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	flash := beego.ReadFromRequest(&c.Controller)
	logs.Debug("parsing parameters")

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				flash.Error("Error: invalid query key/value pair")
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	logs.Debug("get all users")
	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err.Error())
		flash.Error(err.Error())
	} else {
		logs.Debug("len(l)=", len(l))
		c.Data["object_list_len"] = len(l)
		c.Data["object_list"] = l
	}
	c.TplName = "user/index.tpl"

	flash.Store(&c.Controller)
}

// GetAddForm ...
func (c *MyUserController) GetAddForm() {
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "user/create.tpl"
	flash.Store(&c.Controller)
}

// PostAddForm
func (c *MyUserController) PostAddForm() {
	var msg string
	flash := beego.ReadFromRequest(&c.Controller)

	c.TplName = "user/create.tpl"

	// Check XSRF first.
	if !c.CheckXSRFCookie() {
		msg = "XSRF token missing or incorrect."
		c.Data["error"] = msg
		flash.Error(msg)
		flash.Store(&c.Controller)
		return
	}

	// parse form parameters
	var birthday time.Time
	name := c.GetString("name")
	gender := c.GetString("gender")
	t, err := time.Parse("2006-01-02", c.GetString("birthday"))
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		birthday = time.Time{}
		return
	}
	birthday = t

	var user models.User
	user.Name = name
	user.Gender = gender
	user.Birthday = birthday
	id, err := models.AddUser(&user)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	c.Data["id"] = id
	flash = beego.NewFlash()
	flash.Success("Created successful.")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, "/myuser/")
}
