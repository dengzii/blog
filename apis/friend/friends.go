package friend

import (
	"errors"
	"github.com/dengzii/blog/apis/common"
	"github.com/dengzii/blog/models/friend"
	"github.com/kataras/iris/v12/context"
	"strings"
)

func AddFriendsApi(ctx context.Context) (err error) {

	var f *friend.Friend
	err = ctx.ReadJSON(&f)
	if err != nil {
		return
	}
	if len(f.Name) == 0 || len(f.Url) == 0 || !strings.HasPrefix(f.Url, "http") {
		return errors.New("name or url invalid")
	}
	err = friend.AddFriend(f)
	if err != nil {
		return
	}
	_, err = ctx.JSON(common.SuccessResponse(f))
	return
}

func GetFriendsApi(ctx context.Context) (err error) {

	friends := friend.GetFriend()
	_, err = ctx.JSON(common.SuccessResponse(friends))
	return err
}
