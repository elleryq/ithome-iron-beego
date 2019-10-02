// @APIVersion 1.0.0
// @Title ithome-iron-beego API
// @Description ithome-iron-beego API demo.
// @Contact elleryq@gmail.com
package routers

import (
	"strings"

	"github.com/elleryq/ithome-iron-beego/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterMember = func(ctx *context.Context) {
	if strings.HasPrefix(ctx.Input.URL(), "/login") {
		return
	}

	_, ok := ctx.Input.Session("user_id").(int)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Get")
	beego.Router("/login", &controllers.LoginController{}, "post:Post")
	beego.Router("/logout", &controllers.LogoutController{}, "get:Get")
	beego.Router("/post/:id", &controllers.PostController{}, "get:GetOne")

	// beego.InsertFilter("/myuser/*", beego.BeforeRouter, auth.Basic("foobar", "pass"))
	beego.InsertFilter("/myuser/*", beego.BeforeRouter, FilterMember)
	beego.Router("/myuser", &controllers.MyUserController{}, "get:GetAll")
	beego.Router("/myuser/create", &controllers.MyUserController{}, "get:GetAddForm")
	beego.Router("/myuser/create", &controllers.MyUserController{}, "post:PostAddForm")
	beego.AutoRouter(&controllers.UserController{})

	// Automated API Documentation
	ns := beego.NewNamespace(
		"/v1",
		beego.NSNamespace(
			"/user",
			beego.NSInclude(&controllers.UserController{}),
		),
	)
	beego.AddNamespace(ns)
}
