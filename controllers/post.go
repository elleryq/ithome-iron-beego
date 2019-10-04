package controllers

import (
	"encoding/json"
	"errors"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/elleryq/ithome-iron-beego/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/pagination"
)

//  PostController operations for Post
type PostController struct {
	beego.Controller
}

// URLMapping ...
func (c *PostController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// GetAddForm ...
func (c *PostController) GetCreatePostForm() {
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "post/create.tpl"
	flash.Store(&c.Controller)
}

// Post ...
func (c *PostController) PostCreatePostForm() {
	var msg string
	flash := beego.ReadFromRequest(&c.Controller)
	c.TplName = "post/create.tpl"

	// Check XSRF first.
	if !c.CheckXSRFCookie() {
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		msg = "XSRF token missing or incorrect."
		c.Data["error"] = msg
		flash.Error(msg)
		flash.Store(&c.Controller)
		return
	}

	title := c.GetString("title")
	content := c.GetString("content")
	member_id := c.Ctx.Input.Session("user_id").(int)
	member, err := models.GetMemberById(int64(member_id))
	if err != nil {
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		msg = "No such user"
		c.Data["error"] = msg
		flash.Error(msg)
		flash.Store(&c.Controller)
		return
	}

	var v models.Post
	v.Title = title
	v.Content = content
	v.Member = member
	v.PostedAt = time.Now()
	v.ModifiedAt = time.Now()

	id, err := models.AddPost(&v)
	if err != nil {
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return
	}

	c.Data["id"] = id
	flash = beego.NewFlash()
	flash.Success("Post is created successful.")
	flash.Store(&c.Controller)
	c.Ctx.Redirect(302, beego.URLFor("PostController.GetAll"))
}

// GetOne ...
// @Title Get One
// @Description get Post by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Post
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PostController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetPostById(id)
	if err != nil {
		c.Data["notfound"] = true
	} else {
		c.Data["post"] = v
	}
	c.TplName = "post/detail.tpl"
}

// GetAll ...
// @Title Get All
// @Description get Post
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Post
// @Failure 403
// @router / [get]
func (c *PostController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)

	flash := beego.ReadFromRequest(&c.Controller)

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	} else {
		sortby = []string{"PostedAt"}
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	} else {
		order = []string{"desc"}
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	postsPerPage := 10
	postCount, err := models.CountPost()
	if err != nil {
		logs.Error(err.Error())
	}

	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, postCount)
	l, err := models.GetAllPost(query, fields, sortby, order, int64(paginator.Offset()), int64(postsPerPage))
	if err != nil {
		logs.Error(err.Error())
		flash.Error(err.Error())
	} else {
		logs.Debug("len(l)=", len(l))
		c.Data["object_list_len"] = len(l)
		c.Data["object_list"] = l
	}
	c.TplName = "post/index.tpl"
	flash.Store(&c.Controller)
}

// Put ...
// @Title Put
// @Description update the Post
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Post	true		"body for Post content"
// @Success 200 {object} models.Post
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PostController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Post{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdatePostById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Post
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PostController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeletePost(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
