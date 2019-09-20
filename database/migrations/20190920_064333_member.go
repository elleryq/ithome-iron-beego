package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Member_20190920_064333 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Member_20190920_064333{}
	m.Created = "20190920_064333"

	migration.Register("Member_20190920_064333", m)
}

// Run the migrations
func (m *Member_20190920_064333) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE member(`id` int(11) NOT NULL AUTO_INCREMENT,`username` varchar(32) NOT NULL,`password` varchar(32) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Member_20190920_064333) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `member`")
}
