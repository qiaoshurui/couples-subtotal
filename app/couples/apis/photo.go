package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"go.uber.org/zap"
)

type Photo struct {
	api.Api
}

func (p *Photo) AddPhoto(c *gin.Context) {
	img, err := c.FormFile("file")
	if err != nil {
		logger.Error("接收照片失败", zap.Error(err))
		return
	}
	file, err := img.Open()
	if err != nil {
		logger.Error("接收照片失败", zap.Error(err))
		return
	}
	key := "exampleObject"
	//照片上传
	photo := service.Photo{}
	err = photo.UploadTencent(key, file)
	if err != nil {
		logger.Error("照片新增失败", zap.Error(err))
	}
	res.OkWithMessage(c, "照片新增成功", nil)
	//将照片存入数据库
}
