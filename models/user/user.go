package user

import (
	"errors"
	"github.com/dengzii/blog/db"
	"github.com/dengzii/blog/models/base"
	"time"
)

type User struct {
	base.CommonModel
	Name      string `json:"name",gorm:"unique;not null VARCHAR(191)"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Links     string `json:"links"`
	Likes     int32  `json:"likes"`
	Follower  int32  `json:"follower"`
	Following int32  `json:"following"`
	Github    string `json:"github"`
	Views     int32  `json:"views"`
	PassWd    string `json:"-",gorm:"not null"`
}

func GetUser(name string, passwd string) *User {
	var user User
	tab := []interface{}{&User{}}
	db.CreateTable(tab)
	db.Mysql.Where("name = ? AND pass_wd = ?", name, passwd).Find(&user)
	return &user
}

func AddUser(name string, passwd string, email string) (err error) {
	user := &User{
		CommonModel: base.CommonModel{
			CreatedAt: time.Now().Unix(),
		},
		Name:      name,
		Email:     email,
		Avatar:    "",
		Bio:       "",
		Links:     "",
		Likes:     0,
		Follower:  0,
		Following: 0,
		PassWd:    passwd,
	}
	var exist User
	result := db.Mysql.Where("name = ?", name).First(&exist).RowsAffected
	if result != 0 {
		return errors.New("username already exists")
	}
	db.Insert(&user)
	return
}

func GetUserInfo(name string) interface{} {
	var user User
	result := db.Mysql.Where("name = ?", name).First(&user).RowsAffected
	if result == 0 {
		return nil
	}
	return user
}

func View(name string) bool {
	var user User
	result := db.Mysql.Where("name = ?", name).First(&user).RowsAffected
	if result == 0 {
		return false
	}
	user.Views += 1
	return db.Mysql.Model(&user).Update("views", user.Views).RowsAffected > 0
}

func init() {

}
