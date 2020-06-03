package friend

import (
	"errors"
	"github.com/dengzii/blog/db"
	"github.com/jinzhu/gorm"
)

type Friend struct {
	gorm.Model
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Url     string `json:"url"`
	Email   string `json:"emil"`
	Avatar  string `json:"avatar"`
	Alt     string `json:"alt"`
	Display bool   `json:"-"`
}

func AddFriend(f *Friend) error {
	f.Display = false
	if db.Insert(f).RowsAffected == 0 {
		return errors.New("add friend falied")
	}
	return nil
}

func GetFriend() (f []*Friend) {
	db.Mysql.Where("display = ?", 1).Find(&f)
	return
}
