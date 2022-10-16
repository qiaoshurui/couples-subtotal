package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func DynamicRouter(v1 *gin.RouterGroup) {
	user := &apis.Dynamic{}
	{
		v1.POST("/dynamic", user.AddDynamic)          //动态新增
		v1.DELETE("/dynamic", user.DeleteDynamic)     //动态删除
		v1.PUT("/dynamic", user.UpdateDynamic)        //动态修改
		v1.GET("/dynamic/list", user.GetDynamicList)  //查看动态列表
		v1.GET("/dynamic/:id", user.GetDynamicDetail) //查看动态详情
	}
}
