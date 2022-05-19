package admin

import (
	"github.com/go-basic/uuid"
	"mime/multipart"
	"strings"
)

type MainController struct {
	BaseController
}

// Upload @Title 上传图片
// @Description 上传图片
// @router /upload [post]
func (c *MainController) Upload() {
	f, h, err := c.GetFile("file")
	if err != nil {
		panic("文件上传失败")
	}
	fileArr := strings.Split(h.Filename, ".")
	guid := uuid.New()
	var fileName = guid + "." + fileArr[len(fileArr)-1]
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	url := "upload/" + fileName
	err = c.SaveToFile("file", url)
	if err != nil {
		panic(err)
	}
	fileMap := make(map[string]string)
	fileMap["name"] = h.Filename
	fileMap["url"] = "/" + url
	fileMap["uuid"] = guid
	fileMap["params"] = c.GetString("params")
	c.Success(fileMap)
}
