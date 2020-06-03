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
		token := ctx.URLParam("token")
		if token != "dengjianhua520" {
			ctx.StatusCode(iris.StatusForbidden)
			_, _ = ctx.WriteString("403 Forbidden")
			return
		}
	}
	ctx.Header("Allow", "POST, GET, PUT, DELETE, OPTIONS")
	ctx.Header("Vary", "Access-Control-Request-Method")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	ctx.Header("Access-Control-Request-Headers", "Accept,X-Requested-With,Content-Length, Accept-Encoding,X-CSRF-Token,Authorization,token, Content-Type")
	ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type")
	ctx.Header("Access-Control-Request-Method", "*")
	ctx.Next()
}
