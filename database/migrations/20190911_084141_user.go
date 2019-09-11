package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20190911_084141 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20190911_084141{}
	m.Created = "20190911_084141"

	migration.Register("User_20190911_084141", m)
}

// Run the migrations
func (m *User_20190911_084141) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(64) NOT NULL,`gender` varchar(1) NOT NULL,`birthday` datetime NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *User_20190911_084141) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
