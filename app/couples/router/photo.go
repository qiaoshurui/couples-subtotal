package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func PhotoRouter(v1 *gin.RouterGroup) {
	photo := &apis.Photo{}
	{
		v1.POST("/upload/photo", photo.AddPhoto)    //照片上传
		v1.GET("/photo/list", photo.GetPhotoList)   //查看照片列表
		v1.POST("/delete/photo", photo.DeletePhoto) //删除照片
		v1.POST("copy/photo", photo.CopyPhoto)      //复制照片
	}
}
