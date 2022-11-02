package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func PhotoAlbumRouter(v1 *gin.RouterGroup) {
	photo := &apis.PhotoAlbum{}
	{
		v1.POST("/photo-album", photo.AddPhotoAlbum) //相册上传
	}
}
