package article

import (
	"github.com/dengzii/blog/apis/common"
	"github.com/dengzii/blog/models/article"
	"github.com/kataras/iris/v12/context"
)

type tagJson struct {
	Name  string `json:"name"`
	Style int    `json:"style"`
}

func AddTagApi(ctx context.Context) (err error) {
	tagJson := &tagJson{}
	err = ctx.ReadJSON(tagJson)
	tag := article.AddTag(tagJson.Name, tagJson.Style)
	if tag == nil {

	} else {
		_, err = ctx.JSON(common.SuccessResponse(tag))
	}
	return err
}

func GetTagsApi(ctx context.Context) (err error) {

	tags := article.GetTags()
	_, err = ctx.JSON(common.SuccessResponse(tags))
	return
}
