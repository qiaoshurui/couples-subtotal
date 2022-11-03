package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func PhotoAlbumRouter(v1 *gin.RouterGroup) {
	photoAlbum := &apis.PhotoAlbum{}
	{
		v1.POST("upload/photo-album", photoAlbum.AddPhotoAlbum) //相册上传
		v1.GET("/photo-album/list", photoAlbum.GetAlbumList)    //查看相册列表
		v1.POST("/delete/photo-album", photoAlbum.DeleteAlbum)  //相册删除
	}
}
