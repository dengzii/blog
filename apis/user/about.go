package user

import (
	"github.com/dengzii/blog/apis/common"
	"github.com/dengzii/blog/models/user"
	"github.com/kataras/iris/v12/context"
)

func GetAbout(ctx context.Context) (err error) {

	_, err = ctx.JSON(common.SuccessResponse(user.GetAbout()))
	return nil
}

func AddAbout(ctx context.Context) (err error) {

	about := &user.About{}

	err = ctx.ReadJSON(&about)
	println("=>" + about.Content)
	user.AddAbout(about)
	_, err = ctx.JSON(common.SuccessResponse(nil))
	return nil
}
