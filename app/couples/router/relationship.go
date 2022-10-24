package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func RelationshipRouter(v1 *gin.RouterGroup) {
	relationship := &apis.Relationship{}
	{
		v1.GET("/couple-detail/:id", relationship.CoupleDetailDisplay) //情侣关系展示1
		v1.GET("/couple-detail", relationship.CoupleDetailDisplay2)    //情侣关系展示2
		v1.POST("relationship", relationship.CoupleInvitation)         //情侣关系绑定
		v1.DELETE("relationship", relationship.CoupleUnbound)          //情侣关系解绑
	}
}
