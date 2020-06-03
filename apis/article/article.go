package article

import (
	"errors"
	"github.com/dengzii/blog/apis/common"
	"github.com/dengzii/blog/models/article"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"time"
)

func GetArticlesApi(ctx context.Context) (err error) {

	last, err := ctx.URLParamInt64("last")
	var category = ctx.URLParam("category")
	if last == -1 || err != nil {
		last = time.Now().Unix()
	}
	articles := article.GetArticles(last, category, 10)
	responseJson := common.SuccessResponse(articles)
	_, err = ctx.JSON(responseJson)
	return err
}

func GetArchiveApi(ctx context.Context) (err error) {

	archive := article.GetArchive()
	response := common.SuccessResponse(archive)
	_, err = ctx.JSON(response)
	return err
}

func GetArticleApi(ctx context.Context) (err error) {

	id, err := ctx.Params().GetInt("id")
	if err != nil || id <= 0 {
		err = errors.New("bad request")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	article := article.GetArticle(id)
	if article == nil {
		err = errors.New("article not found")
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	responseJson := common.SuccessResponse(article)
	_, err = ctx.JSON(responseJson)
	return
}

func AddArticleApi(ctx context.Context) (err error) {

	newArticle := &article.Article{}
	err = ctx.ReadJSON(newArticle)

	invalid := len(newArticle.Title) == 0 || len(newArticle.Content) == 0 || newArticle.AuthorId == 0
	if err != nil || invalid {
		return
	}
	err, result := article.AddArticle(newArticle)
	if err != nil {
		return
	}
	_, err = ctx.JSON(common.SuccessResponse(result))
	return
}

func ViewArticleApi(ctx context.Context) (err error) {

	articleId, err := ctx.Params().GetInt("id")
	if err != nil || articleId <= 0 {
		err = errors.New("bad request")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = article.ViewArticle(articleId)
	if err != nil {
		return
	}
	_, err = ctx.JSON(common.SuccessResponse(""))
	return
}

func LikeArticleApi(ctx context.Context) (err error) {
	articleId, err := ctx.Params().GetInt("id")
	if err != nil || articleId <= 0 {
		err = errors.New("bad request")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	err = article.LikeArticle(articleId)
	if err != nil {
		return
	}
	_, err = ctx.JSON(common.SuccessResponse("thank you."))
	return err
}

func CommentArticleApi(ctx context.Context) (err error) {

	return err
}
