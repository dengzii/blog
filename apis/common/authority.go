package common

import (
	"github.com/kataras/iris/v12/context"
)

var authorityUrl map[string]string

func init() {

}

func AuthorityController(ctx context.Context) {

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
