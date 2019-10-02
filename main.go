package main

import (
	"fmt"
	"os"

	"github.com/elleryq/ithome-iron-beego/checks"
	"github.com/elleryq/ithome-iron-beego/controllers"
	_ "github.com/elleryq/ithome-iron-beego/global"
	_ "github.com/elleryq/ithome-iron-beego/routers"

	// _ "github.com/elleryq/ithome-iron-beego/tasks"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql" // import your required driver
)

var DBARGS = struct {
	Driver string
	Source string
	Debug  string
}{
	os.Getenv("ORM_DRIVER"),
	os.Getenv("ORM_SOURCE"),
	os.Getenv("ORM_DEBUG"),
}

func init() {
	Debug, _ := orm.StrTo(DBARGS.Debug).Bool()

	fmt.Printf("Debug=%v\n", Debug)
	if Debug {
		fmt.Printf("ORM config=%v\n", DBARGS)
	}
	if DBARGS.Driver == "" || DBARGS.Source == "" {
		fmt.Println("Please set ORM_DRIVER/ORM_SOURCE")
		os.Exit(2)
	}

	// set default database
	orm.RegisterDataBase("default", DBARGS.Driver, DBARGS.Source, 30)

	//
	controllers.InitApp()

	// toolbox
	toolbox.AddHealthCheck("database", &checks.DatabaseCheck{})

	// static
	beego.SetStaticPath("/swagger", "swagger")
}

func main() {
	beego.Run()
}
