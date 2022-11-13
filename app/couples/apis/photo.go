package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
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
// @Router /api/v1/upload/photo [post]
func (p *Photo) AddPhoto(c *gin.Context) {
	albumId := c.PostForm("albumId")
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

// GetPhotoList
// @Tags Photo
// @Summary 获取照片列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dto.PhotoList true "页码, 每页大小,选填：相册id"
// @Success 200 {object} dto.PhotoListRes "分页获取照片列表,返回包括列表,总数"
// @Router /api/v1/photo/list [get]
func (p *Photo) GetPhotoList(c *gin.Context) {
	page, size := getPageInfo(c)
	albumId := c.Query("album-id")
	parseInt, _ := strconv.ParseInt(albumId, 10, 64)
	data := &dto.PhotoList{
		Page:     page,
		PageSize: size,
		AlbumId:  parseInt,
	}
	photo := service.Photo{}
	photoList, err := photo.GetPhotoList(data)
	if err != nil {
		logger.Error("照片列表查找失败", zap.Error(err))
	}
	res.OkWithMessage(c, "照片列表展示成功", photoList)
}

// DeletePhoto
// @Tags Photo
// @Summary 照片删除
// @Security ApiKeyAuth
// @Produce application/json
// @Param ids body []int64 true "照片id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"照片删除成功"}"
// @Router /api/v1/delete/photo [post]
func (p *Photo) DeletePhoto(c *gin.Context) {
	var ids []int64
	if err := c.ShouldBindJSON(&ids); err != nil {
		logger.Error("照片删除请求参数有误", zap.Error(err))
	}
	photo := service.Photo{}
	err := photo.DeletePhoto(ids)
	if err != nil {
		logger.Error("数据库照片删除失败", zap.Error(err))
	}
	//删除cos照片
	err = photo.DeleteCosRecord(ids)
	if err != nil {
		logger.Error("腾讯云cos照片删除失败", zap.Error(err))
	}
	res.OkWithMessage(c, "照片删除成功", nil)
}
func (p *Photo) CopyPhoto(c *gin.Context) {
	var ids []int64
	if err := c.ShouldBindJSON(&ids); err != nil {
		logger.Error("照片复制请求参数有误", zap.Error(err))
	}
	//photo := service.Photo{}
	//复制cos照片
	//photo.CopyCosRecord()

}
