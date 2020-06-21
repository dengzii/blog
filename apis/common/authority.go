package common

import (
	"fmt"
	"github.com/dengzii/blog/db"
	"github.com/dengzii/blog/models/user"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/spf13/cast"
	"strings"
	"time"
)

var noAuthorityUrl []string

const (
	TokenExpiredSec = 86400
)

func init() {
	noAuthorityUrl = []string{
		//"/about",
		//"/article",
		//"/tag",
		//"/category",
		"/friend",
		"/user/views",
		"/user",
	}
}

func AuthorityController(ctx context.Context) {

	needAuthor := true
	accessDenied := false

	for i := range noAuthorityUrl {
		if strings.Contains(ctx.Path(), noAuthorityUrl[i]) {
			needAuthor = false
			break
		}
	}

	if needAuthor && ctx.Method() == iris.MethodPut {
		var token = ctx.GetCookie("token")
		accessDenied = true
		if len(token) > 0 {
			var userName = ctx.Params().GetString("username")
			var log user.Log
			var row = db.Mysql.Where("token = ? AND user_name = ?", token, userName).First(&log).RowsAffected
			fmt.Println(cast.ToString(time.Now().Unix() - log.CreatedAt))
			if row > 0 && (time.Now().Unix()-log.CreatedAt) < TokenExpiredSec {
				accessDenied = false
			}
		}
	}
	if accessDenied {
		ctx.StatusCode(iris.StatusForbidden)
		_, _ = ctx.WriteString("Access Denied!")
		return
	}
	ctx.Next()
}
