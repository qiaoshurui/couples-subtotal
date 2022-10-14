package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func DynamicRouter(v1 *gin.RouterGroup) {
	user := &apis.Dynamic{}
	{
		v1.POST("/addDynamic", user.AddDynamic)         //动态新增
		v1.DELETE("/deleteDynamic", user.DeleteDynamic) //动态删除
		v1.PUT("/updateDynamic", user.UpdateDynamic)    //动态修改
		v1.GET("/getDynamic", user.GetDynamicList)      //动态查找
	}
}
