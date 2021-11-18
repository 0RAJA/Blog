package routers

import (
	"Blog/global"
	"Blog/internal/service"
	"Blog/pkg/app"
	"Blog/pkg/convert"
	"Blog/pkg/errcode"
	"Blog/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// UpLoadFile
// @Tags File
// @Summary 上传图片
// @Produce  json
// @Param file formData string true "文件名称"
// @Param type body int true "1为图片"
// @Success 200 {string} string "URL"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /upload/file [post]
func (u Upload) UpLoadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file") //获取文件信息
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType := convert.StrTo(c.PostForm("type")).MustInt() //获取文件类型
	if fileHeader == nil || fileType == 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader) //调用svc进行图片保存
	if err != nil {
		global.Logger.Infof("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{"file_access_url": fileInfo.AccessUrl}) //返回文件链接给前端
}
