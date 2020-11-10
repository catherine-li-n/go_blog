package api

import (
	"github.com/catherine.li/go_blog/global"
	"github.com/catherine.li/go_blog/internal/service"
	"github.com/catherine.li/go_blog/pkg/app"
	"github.com/catherine.li/go_blog/pkg/convert"
	"github.com/catherine.li/go_blog/pkg/errcode"
	"github.com/catherine.li/go_blog/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 通过c.Request.FormFile读取入参file字段的上传文件信息，并把入参type字段作为上传文件类型的确立依据
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
