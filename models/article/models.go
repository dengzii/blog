package article

import (
	"github.com/dengzii/blog/models/base"
	"github.com/dengzii/blog/tools"
	"github.com/jinzhu/gorm"
	"time"
)

type ArticleBase struct {
	base.CommonModel
	CID          string `json:"-"`
	Title        string `json:"title,omitempty" validate:"required"`
	AuthorId     uint   `json:"author_id"  validate:"required"`
	AuthorName   string `json:"author_name,omitempty default:dengzi"`
	Description  string `json:"description,omitempty"`
	TagName      string `json:"tag_name,omitempty" validate:"required"`
	CategoryName string `json:"category_name,omitempty" validate:"required"`
	Likes        uint   `json:"likes,omitempty default:0"`
	Views        uint   `json:"views,omitempty default:0"`
}

type Article struct {
	ArticleBase

	Comments uint     `json:"comments,omitempty default:0"`
	Tag      Tag      `json:"tag,omitempty" gorm:"ForeignKey:TagId"`
	Category Category `json:"category,omitempty" gorm:"ForeignKey:CategoryId"`
	Content  string   `json:"content,omitempty" gorm:"type:TEXT"`
	Display  bool     `json:"-" gorm:"default:true"`
}

func (that *Article) BeforeCreate(scope *gorm.Scope) error {
	tools.Log("Ready to create article,", that)
	that.CreatedAt = time.Now().Unix()
	that.UpdatedAt = that.CreatedAt
	return nil
}

func (that *Article) AfterCreate(scope *gorm.Scope) error {

	//tools.Log("Ready to create article,", article)
	//scope.SetColumn("ID", time.Now())
	that.CreatedAt = time.Now().Unix()
	that.UpdatedAt = that.CreatedAt
	return nil
}

func (that *Article) BeforeDelete(scope *gorm.Scope) error {
	//tools.Log("Ready to create article,", article)
	//scope.SetColumn("ID", time.Now())
	return nil
}

type Archive struct {
	base.CommonModel
	ArticleId uint   `json:"article_id"`
	Title     string `json:"title"`
}

type Category struct {
	base.CommonModel
	Name         string `json:"name"`
	ArticleCount int    `json:"article_count"`
	Display      bool   `json:"-"`
}

type Tag struct {
	base.CommonModel
	ClassId      uint   `json:"class_id"`
	Name         string `json:"name"`
	ArticleCount int    `json:"article_count"`
	Display      bool   `json:"-"`
	Style        int    `json:"-"`
}

func (that *Article) toArticleBase() (articleBase *ArticleBase) {
	articleBase = &ArticleBase{
		Title:        that.Title,
		AuthorName:   that.AuthorName,
		Description:  that.Description,
		TagName:      that.TagName,
		CategoryName: that.CategoryName,
		Likes:        that.Likes,
		Views:        that.Views,
	}
	articleBase.CreatedAt = that.CreatedAt
	articleBase.UpdatedAt = that.UpdatedAt
	return
}
