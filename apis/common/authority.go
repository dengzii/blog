package common

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"strings"
)

var noAuthorityUrl []string

func init() {
	noAuthorityUrl = []string{
		"/friend",
		"/user/views",
	}
}

func AuthorityController(ctx context.Context) {

	needAuthor := true
	for i := range noAuthorityUrl {
		if strings.Contains(ctx.Path(), noAuthorityUrl[i]) {
			needAuthor = false
			break
		}
	}

	if needAuthor && ctx.Method() == iris.MethodPut {
		//token := ctx.URLParam("token")
		_, _ = ctx.WriteString("403")
		return
	}
	ctx.Next()
}
