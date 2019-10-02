package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Post_20191001_071922 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Post_20191001_071922{}
	m.Created = "20191001_071922"

	migration.Register("Post_20191001_071922", m)
}

// Run the migrations
func (m *Post_20191001_071922) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE post(`id` int(11) NOT NULL AUTO_INCREMENT,`member_id` int(11) DEFAULT NULL,`title` varchar(128) NOT NULL,`content` text NULL,`posted_at` datetime NOT NULL,`modified_at` datetime NOT NULL,PRIMARY KEY (`id`),FOREIGN KEY(member_id) REFERENCES member(id))")
}

// Reverse the migrations
func (m *Post_20191001_071922) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `post`")
}
