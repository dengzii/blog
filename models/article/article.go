package article

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/dengzii/blog/db"
	"github.com/dengzii/blog/models/base"
	user2 "github.com/dengzii/blog/models/user"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

func AddArticle(newArticle *Article) (err error, articleBase *ArticleBase) {

	index := strings.Index(newArticle.Content, "\n")
	var desc string
	if index != -1 {
		desc = newArticle.Content[0:index]
	} else {
		desc = newArticle.Content
	}
	newArticle.Description = desc

	h := md5.New()
	h.Write([]byte(newArticle.Title))
	newArticle.CID = hex.EncodeToString(h.Sum(nil))

	newArticle.CreatedAt = time.Now().Unix()
	newArticle.UpdatedAt = newArticle.CreatedAt

	var user user2.User
	findUsr := db.Mysql.Where("id = ?", newArticle.AuthorId).First(&user).RowsAffected
	if findUsr == 0 {
		return errors.New("no such user"), nil
	}
	newArticle.AuthorName = user.Name
	if db.Insert(newArticle).RowsAffected != 0 {
		db.Insert(&Archive{
			CommonModel: base.CommonModel{
				CreatedAt: newArticle.CreatedAt,
				UpdatedAt: newArticle.CreatedAt,
			},
			ArticleId: newArticle.ID,
			Title:     newArticle.Title,
		})
		err = nil
	} else {
		err = errors.New("create article filed")
	}
	articleBase = newArticle.toArticleBase()
	return
}

func GetArticles(from int64, category string, count int) (articles []*ArticleBase) {
	var article []*Article

	var query *gorm.DB
	query = db.Mysql.Order("updated_at desc")

	if len(category) > 0 && category != "All" {
		query = query.Where("category_name = ?", category)
	}

	query.
		Where("updated_at < ?", from).
		Limit(count).
		Find(&article)

	articles = make([]*ArticleBase, len(article))
	for i, v := range article {
		articles[i] = v.toArticleBase()
	}
	return articles
}

func GetArticle(id int, view bool) *Article {
	var article Article
	db.Mysql.Where("id = ?", id).Attrs(nil).FirstOrInit(&article)
	if len(article.Title) == 0 {
		return nil
	}
	if view {
		article.Views += 1
		db.Mysql.Model(&article).Update("views", article.Views)
	}
	return &article
}

func ViewArticle(id int) (err error) {
	var article Article
	db.Mysql.Where("id = ?", id).First(&article)
	if len(article.Title) == 0 {
		err = errors.New("article not found")
		return
	}
	article.Views += 1
	db.Mysql.Model(&article).Update("views", article.Views)
	return
}

func LikeArticle(id int) (err error) {
	var article Article
	db.Mysql.Where("id = ?", id).First(&article)
	if len(article.Title) == 0 {
		err = errors.New("article not found")
		return
	}
	article.Likes += 1
	db.Mysql.Model(&article).Update("likes", article.Likes)
	return
}

func GetArchive() (archive []*Archive) {
	db.Mysql.Order("created_at desc").Find(&archive)
	return
}

func DelArticle(id int) {

}

func UpdateArticle() {

}

func CommentArticle() {

}
