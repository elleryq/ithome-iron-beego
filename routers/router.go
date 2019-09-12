package routers

import (
	"my/hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/about", &controllers.AboutController{})
	beego.Router("/myuser", &controllers.MyUserController{}, "get:GetAll")
	beego.AutoRouter(&controllers.UserController{})
}
