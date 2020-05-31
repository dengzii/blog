package apis

import (
	"fmt"
	"github.com/dengzii/blog/apis/common"
	"github.com/kataras/iris/v12/context"
	"strings"
)

func StaticFileApi(context context.Context) error {

	filePath := context.Params().Get("file")
	userName := context.Params().GetString("username")

	if strings.HasPrefix(filePath, ".") {
		common.ReturnNotFound(context)
		return nil
	}
	path := fmt.Sprintf("./statics/%s/%s", userName, filePath)
	err := context.ServeFile(path, false)
	if err != nil {
		common.ReturnNotFound(context)
	}
	return nil
}

func UploadFileApi(context context.Context) error {

	common.ReturnNotFound(context)
	return nil
}
