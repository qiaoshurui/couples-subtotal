package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"go.uber.org/zap"
	"strconv"
)

type Photo struct {
	api.Api
}

// AddPhoto
// @Tags Photo
// @Summary 照片上传
// @Security ApiKeyAuth
// @Produce application/json
// @Param albumId formData int64 true "相册id"
// @Param file formData file true "照片上传"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"照片上传成功"}"
// @Router /api/v1/photo [post]
func (p *Photo) AddPhoto(c *gin.Context) {
	albumId := c.PostForm("albumId")
	fmt.Println(albumId)
	img, err := c.FormFile("file")
	if err != nil {
		logger.Error("接收文件失败", zap.Error(err))
		return
	}
	file, err := img.Open()
	if err != nil {
		logger.Error("文件打开失败", zap.Error(err))
		return
	}
	imgName := img.Filename
	photo := service.Photo{}
	parseInt, err := strconv.ParseInt(albumId, 10, 64)
	//照片上传到cos
	imgUrl, err := photo.UploadTencent(imgName, parseInt, file)
	if err != nil {
		logger.Error("照片上传到cos失败", zap.Error(err))
	}
	//将照片存入数据库
	err = photo.AddPhoto(imgUrl, parseInt)
	if err != nil {
		logger.Error("照片上传到数据库失败", zap.Error(err))
	}
	res.OkWithMessage(c, "照片新增成功", nil)
}
