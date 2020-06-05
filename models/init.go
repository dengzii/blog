package models

import (
	"github.com/dengzii/blog/db"
	"github.com/dengzii/blog/models/article"
	"github.com/dengzii/blog/models/friend"
	"github.com/dengzii/blog/models/user"
)

func Init() {
	tableArticle := &article.Article{}
	tableUser := &user.User{}
	tableTag := &article.Tag{}
	tableClass := &article.Category{}
	tableFriend := &friend.Friend{}
	tableArchive := &article.Archive{}

	tab := []interface{}{
		tableUser, tableTag, tableClass, tableFriend, tableArticle, tableArchive,
		&Comment{}, &user.About{}, &user.Log{},
	}

	db.CreateTables(tab)
}
