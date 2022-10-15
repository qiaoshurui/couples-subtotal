package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"go.uber.org/zap"
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
	fmt.Printf("请求路径：%s 请求参数，id：%v", c.Request.URL, id)
	if err != nil {
		logger.Error("情侣关系展示页面请求参数有误", zap.Error(err))
		return
	}
	relationship := service.Relationship{}
	getRelationship, err := relationship.GetRelationship(id)
	if err != nil {
		logger.Error("情侣关系页面展示失败", zap.Error(err))
	}
	res.OkWithMessage(c, "情侣关系页面展示成功", getRelationship)

}

func (r Relationship) CoupleDetailDisplay2(c *gin.Context) {
	//参数id
	userId := c.Query("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	fmt.Printf("请求路径：%s 请求参数，id：%v", c.Request.URL, id)
	if err != nil {
		logger.Error("情侣关系展示页面请求参数有误", zap.Error(err))
		return
	}
	relationship := service.Relationship{}
	getRelationship, err := relationship.GetRelationship(id)
	if err != nil {
		logger.Error("情侣关系页面展示失败", zap.Error(err))
	}
	res.OkWithMessage(c, "情侣关系页面展示成功", getRelationship)

}
