package controllers

import (
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

// GetEditPostForm ...
func (c *PostController) GetEditPostForm() {
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetPostById(id)
	if err != nil {
		c.Data["notfound"] = true
	} else {
		c.Data["post"] = v
	}
	c.TplName = "post/edit.tpl"
	flash.Store(&c.Controller)
}

// Post ...
func (c *PostController) PostEditPostForm() {
	var msg string
	var id int64
	var err error
	var title, content string
	var post *models.Post

	flash := beego.ReadFromRequest(&c.Controller)
	c.TplName = "post/edit.tpl"

	// Check XSRF first.
	if !c.CheckXSRFCookie() {
		msg = "XSRF token missing or incorrect."
		goto ERROR
	}

	id, err = c.GetInt64("id")
	if err != nil {
		msg = "Bad form"
		goto ERROR
	}

	title = c.GetString("title")
	content = c.GetString("content")
	post, err = models.GetPostById(id)
	if err != nil {
		msg = "No such post"
		goto ERROR
	}

	post.Title = title
	post.Content = content
	post.ModifiedAt = time.Now()

	err = models.UpdatePostById(post)
	if err != nil {
		msg = err.Error()
		goto ERROR
	}

	c.Data["id"] = id
	flash = beego.NewFlash()
	flash.Success("Post is updated successful.")
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

// GetCreatePostForm ...
func (c *PostController) GetCreatePostForm() {
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "post/create.tpl"
	flash.Store(&c.Controller)
}

// Post ...
func (c *PostController) PostCreatePostForm() {
	var msg string
	var id int64
	var title string
	var content string
	var member_id int
	var member *models.Member
	var v models.Post
	var err error

	flash := beego.ReadFromRequest(&c.Controller)
	c.TplName = "post/create.tpl"

	// Check XSRF first.
	if !c.CheckXSRFCookie() {
		msg = "XSRF token missing or incorrect."
		goto ERROR
	}

	title = c.GetString("title")
	content = c.GetString("content")
	member_id = c.Ctx.Input.Session("user_id").(int)
	member, err = models.GetMemberById(int64(member_id))
	if err != nil {
		msg = "No such user"
		goto ERROR
	}

	v.Title = title
	v.Content = content
	v.Member = member
	v.PostedAt = time.Now()
	v.ModifiedAt = time.Now()

	id, err = models.AddPost(&v)
	if err != nil {
		msg = err.Error()
		goto ERROR
	}

	c.Data["id"] = id
	flash = beego.NewFlash()
	flash.Success("Post is created successful.")
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
	var err error
	var l []interface{}
	var postsPerPage int = 10
	var postCount int64 = 0
	var paginator *pagination.Paginator

	c.TplName = "post/index.tpl"
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
				err = errors.New("Error: invalid query key/value pair")
				goto ERROR
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	postCount, err = models.CountPost()
	if err != nil {
		goto ERROR
	}

	paginator = pagination.SetPaginator(c.Ctx, postsPerPage, postCount)
	l, err = models.GetAllPost(query, fields, sortby, order, int64(paginator.Offset()), int64(postsPerPage))
	if err != nil {
		goto ERROR
	}

	logs.Debug("len(l)=", len(l))
	c.Data["object_list_len"] = len(l)
	c.Data["object_list"] = l
	goto FINAL

ERROR:
	logs.Error(err.Error())
	flash.Error(err.Error())

FINAL:
	flash.Store(&c.Controller)
}

// GetDeletePostForm ...
func (c *PostController) GetDeletePostForm() {
	flash := beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetPostById(id)
	if err != nil {
		c.Data["notfound"] = true
	} else {
		c.Data["post"] = v
	}
	c.TplName = "post/delete.tpl"
	flash.Store(&c.Controller)
}

func (c *PostController) PostDeletePostForm() {
	var msg string
	var err error
	var id int64

	flash := beego.ReadFromRequest(&c.Controller)
	c.TplName = "post/delete.tpl"

	// Check XSRF first.
	if !c.CheckXSRFCookie() {
		msg = "XSRF token missing or incorrect."
		goto ERROR
	}

	id, err = c.GetInt64("id")
	if err != nil {
		msg = "Bad form"
		goto ERROR
	}

	err = models.DeletePost(id)
	if err != nil {
		msg = err.Error()
		goto ERROR
	}
	goto SUCCESS

SUCCESS:
	c.Data["id"] = id
	flash = beego.NewFlash()
	flash.Success("Post is updated successful.")
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
