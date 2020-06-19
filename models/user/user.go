package user

import (
	"errors"
	"github.com/dengzii/blog/db"
	"github.com/dengzii/blog/models/base"
	"math/rand"
	"time"
)

type User struct {
	base.CommonModel
	Name      string `json:"name",gorm:"unique;not null VARCHAR(191)"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	SiteName  string `json:"site_name"`
	Bio       string `json:"bio"`
	Links     string `json:"links"`
	Likes     int32  `json:"likes"`
	Follower  int32  `json:"follower"`
	Following int32  `json:"following"`
	Github    string `json:"github"`
	Views     int32  `json:"views"`
	PassWd    string `json:"-",gorm:"not null"`
}

type Log struct {
	base.CommonModel
	UserId   uint
	UserName string
	Token    string
}

func GetUser(name string, passwd string) (user User, token string) {
	var result = db.Mysql.Where("name = ? AND pass_wd = ?", name, passwd).Find(&user).RowsAffected
	if result > 0 {
		token = getRandomString(32)
		db.Insert(&Log{
			CommonModel: base.CommonModel{
				CreatedAt: time.Now().Unix(),
			},
			UserName: user.Name,
			UserId:   user.ID,
			Token:    token,
		})
	}
	return
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

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
