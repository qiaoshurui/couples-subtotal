package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func PhotoRouter(v1 *gin.RouterGroup) {
	photo := &apis.Photo{}
	{
		v1.POST("/photo", photo.AddPhoto)            //照片上传
		v1.POST("/photo-album", photo.AddPhotoAlbum) //相册上传
	}
}
