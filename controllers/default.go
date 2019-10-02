package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/elleryq/ithome-iron-beego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)

	sortby = []string{"PostedAt"}
	order = []string{"desc"}
	postsPerPage := 10
	postCount, err := models.CountPost()
	if err != nil {
	}
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, postCount)

	// fetch the next 20 posts
	posts, err := models.GetAllPost(query, fields, sortby, order, int64(paginator.Offset()), int64(postsPerPage))
	if err != nil {
		logs.Error(err.Error())
	}
	c.Data["posts"] = posts
	/*
		post := posts[0].(models.Post)
		logs.Debug("1.", post.Title)
		logs.Debug("2.", post.Member.Username)
	*/
	c.TplName = "index.tpl"
}
