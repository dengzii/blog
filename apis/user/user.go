package user

import (
	"github.com/dengzii/blog/apis/common"
	user2 "github.com/dengzii/blog/models/user"
	"github.com/kataras/iris/v12/context"
	"net/http"
	"time"
)

type loginJson struct {
	Name    string `json:"username"  validate:"required"`
	Passwd  string `json:"password" validate:"required"`
	Captcha string `json:"captcha" validate:"required"`
}

type registerJson struct {
	Name    string `json:"name"`
	Passwd  string `json:"passwd"`
	Email   string `json:"email"`
	Captcha string `json:"captcha"`
}

type userIdJson struct {
	Name string `json:"name"`
}

func LoginApi(ctx context.Context) (err error) {

	var requestUser loginJson

	err = ctx.ReadJSON(&requestUser)
	user, token := user2.GetUser(requestUser.Name, requestUser.Passwd)

	if len(token) == 0 {
		_, err = ctx.JSON(common.ErrorEmptyResponse("login failed"))
		return
	}
	ctx.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, 1),
	})
	_, err = ctx.JSON(common.SuccessResponse(user))
	return err
}

func RegisterApi(ctx context.Context) (err error) {

	var user registerJson
	err = ctx.ReadJSON(&user)
	if err != nil {
		return err
	}
	err = user2.AddUser(user.Name, user.Passwd, user.Email)
	if err != nil {
		_, err = ctx.JSON(common.SuccessResponse(`welcome ` + user.Name))
	}
	return
}

func GetUserInfoApi(ctx context.Context) (err error) {
	name := ctx.Params().GetString("username")
	profile := user2.GetUserInfo(name)
	_, err = ctx.JSON(common.SuccessResponse(profile))
	return err
}

func ViewSiteApi(ctx context.Context) (err error) {
	name := ctx.Params().GetString("username")
	success := user2.View(name)
	if success {
		_, err = ctx.JSON(common.SuccessResponse(nil))
	} else {
		_, err = ctx.JSON(common.ErrorEmptyResponse("failed."))
	}
	return err
}
