package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"strconv"
)

type Relationship struct {
	api.Api
}

// CoupleDetailDisplay 情侣关系展示页面
func (r Relationship) CoupleDetailDisplay(c *gin.Context) {
	//参数id
	userId := c.Param("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		//输出日志
		//返回响应
		return
	}
	relationship := service.Relationship{}
	getRelationship, err := relationship.GetRelationship(id)
	if err != nil {
		//输出日志
		//返回响应
	}
	res.OkWithMessage(c, "情侣关系页面展示成功", getRelationship)
}
