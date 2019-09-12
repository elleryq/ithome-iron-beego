package controllers

import (
	"errors"
	"my/hello/models"
	"strings"

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

	c.Data["has_error"] = false
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
				c.Data["has_error"] = true
				c.Data["error"] = errors.New("Error: invalid query key/value pair")
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
		// c.Data["error"] = err.Error()
		c.Data["has_error"] = true
		c.Data["error"] = "errors"
	} else {
		logs.Debug("len(l)=", len(l))
		c.Data["object_list_len"] = len(l)
		c.Data["object_list"] = l
	}
	c.TplName = "user/index.tpl"
}
