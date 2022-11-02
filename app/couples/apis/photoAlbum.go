package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"github.com/qiaoshurui/couples-subtotal/common/utils"
	"go.uber.org/zap"
)

type PhotoAlbum struct {
	api.Api
}

// AddPhotoAlbum
// @Tags PhotoAlbum
// @Summary 相册新增
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body dto.AddPhotoAlbum true "相册名称，相册类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"相册创建成功"}"
// @Router /api/v1/photo-album  [post]
func (p *PhotoAlbum) AddPhotoAlbum(c *gin.Context) {
	var album dto.AddPhotoAlbum
	if err := c.ShouldBindJSON(&album); err != nil {
		logger.Error("相册上传请求参数有误", zap.Error(err))
		return
	}
	//相册的uuid
	uuid := utils.RegCodeCreat()
	photoAlbum := service.PhotoAlbum{}
	//cos新增相册
	key, err := photoAlbum.UploadTencent(album.Type, uuid)
	//数据库新增相册
	err = photoAlbum.AddPhotoAlbum(&album, key, uuid)
	if err != nil {
		logger.Error("相册新增失败", zap.Error(err))
	}
	res.OkWithMessage(c, "相册新增成功", nil)
}
