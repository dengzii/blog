package article

import (
	"github.com/dengzii/blog/apis/common"
	"github.com/dengzii/blog/models/article"
	"github.com/kataras/iris/v12/context"
)

type categoryJson struct {
	Name string
}

func AddCategoryApi(ctx context.Context) (err error) {
	categoryJson := &categoryJson{}
	err = ctx.ReadJSON(categoryJson)
	category := article.AddCategory(categoryJson.Name)
	if category == nil {

	} else {
		_, err = ctx.JSON(common.SuccessResponse(category))
	}
	return err
}

func GetCategoriesApi(ctx context.Context) (err error) {

	tags := article.GetCategories()
	_, err = ctx.JSON(common.SuccessResponse(tags))
	return
}
