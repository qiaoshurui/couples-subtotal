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
	"strconv"
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
// @Router /api/v1/upload/photo-album  [post]
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

// GetAlbumList
// @Tags PhotoAlbum
// @Summary 获取相册列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dto.AlbumListReq true "页码, 每页大小,相册类型"
// @Success 200 {object} dto.AlbumListRes "分页获取相册列表成功"
// @Router /api/v1/dynamic/list [get]
func (p *PhotoAlbum) GetAlbumList(c *gin.Context) {
	page, size := getPageInfo(c)
	genre := c.Query("type")
	parseInt, _ := strconv.ParseInt(genre, 10, 64)
	data := &dto.AlbumListReq{
		Page:     page,
		PageSize: size,
		Type:     int8(parseInt),
	}
	album := service.PhotoAlbum{}
	albumList, err := album.GetAlbumList(data)
	if err != nil {
		logger.Error("查找相册列表失败", zap.Error(err))
	}
	res.OkWithMessage(c, "查找相册列表成功", albumList)
}

// DeleteAlbum
// @Tags PhotoAlbum
// @Summary 相册删除
// @Security ApiKeyAuth
// @Produce application/json
// @Param ids body []int64 true "相册id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"相册及相册照片删除成功"}"
// @Router /api/v1/delete/photo-album [post]
func (p *PhotoAlbum) DeleteAlbum(c *gin.Context) {
	var ids []int64
	if err := c.ShouldBindJSON(&ids); err != nil {
		logger.Error("相册删除请求参数有误", zap.Error(err))
	}
	photoAlbum := service.PhotoAlbum{}
	err := photoAlbum.DeleteAlbum(ids)
	if err != nil {
		logger.Error("数据库照片删除失败", zap.Error(err))
	}
	//删除cos相册
	err = photoAlbum.DeleteCosRecord(ids)
	if err != nil {
		logger.Error("腾讯云cos相册删除失败", zap.Error(err))
	}
	res.OkWithMessage(c, "相册及相册照片删除成功", nil)
}
