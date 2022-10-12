package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func RelationshipRouter(v1 *gin.RouterGroup) {
	relationship := &apis.Relationship{}
	{
		v1.POST("/display", relationship.Display) //情侣关系展示
	}
}
