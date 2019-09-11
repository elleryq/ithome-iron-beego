package main

import (
	"fmt"
	_ "my/hello/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
}

func main() {
	// Create
	/*
		userNotAdded := models.User{Name: "John Doe", Birthday: time.Now(), Gender: "M"}
		userid, err := models.AddUser(&userNotAdded)
		if err != nil {
			panic(err)
		}
		fmt.Println(userid)
	*/

	// Query
	/*
		query := map[string]string{}
		fields := []string{}
		sortby := []string{}
		order := []string{}
		users, err := models.GetAllUser(query, fields, sortby, order, 0, -1)
		if err != nil {
			panic(err)
		}
		fmt.Println(users)
	*/

	// Retrieve single object
	/*
		user, err := models.GetUserById(2)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
	*/

	// Update
	/*
		user.Name = "Mary Jane"
		user.Gender = "F"
		err = models.UpdateUserById(user)
		if err != nil {
			panic(err)
		}
		// Retrieve again to verify
		user, err = models.GetUserById(2)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
	*/

	// Delete
	/*
		err := models.DeleteUser(2)
		if err != nil {
			panic(err)
		}
	*/

	beego.Run()
}
