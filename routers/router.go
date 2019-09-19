package routers

import (
	"my/hello/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/auth"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})

	beego.InsertFilter("/myuser/*", beego.BeforeRouter, auth.Basic("foobar", "pass"))
	beego.Router("/myuser", &controllers.MyUserController{}, "get:GetAll")
	beego.Router("/myuser/create", &controllers.MyUserController{}, "get:GetAddForm")
	beego.Router("/myuser/create", &controllers.MyUserController{}, "post:PostAddForm")
	beego.AutoRouter(&controllers.UserController{})
}
