package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"go.uber.org/zap"
	"strconv"
)

type Relationship struct {
	api.Api
}

// CoupleDetailDisplay
// @Tags Relationship
// @Summary 情侣关系展示页面
// @Security ApiKeyAuth
// @Produce application/json
// @Param "id" query int64 true "ID"
// @Success 200 {object} dto.CouplesInfo count "情侣关系展示页面获取成功"
// @Router /api/v1/couple-detail/:id [get]
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

// CoupleDetailDisplay2
// @Tags Relationship
// @Summary 情侣关系展示页面
// @Security ApiKeyAuth
// @Produce application/json
// @Param id query int64 true "ID"
// @Success 200 {object} dto.CouplesInfo count "情侣关系展示页面获取成功"
// @Router /api/v1/couple-detail [get]
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

// CoupleInvitation
// @Tags Relationship
// @Summary 情侣关系绑定（通过链接进行关系绑定）
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dto.RelationshipBinding true "加密后的注册码，用户id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"情侣关系绑定成功"}"
// @Router /api/v1/relationship [post]
func (r Relationship) CoupleInvitation(c *gin.Context) {
	var bindingParam dto.RelationshipBinding
	if err := c.ShouldBindJSON(&bindingParam); err != nil {
		logger.Error("关系绑定请求参数有误", zap.Error(err))
		return
	}
	relationship := service.Relationship{}
	err := relationship.RelationBinding(&bindingParam)
	if err != nil {
		logger.Error("情侣关系绑定失败", zap.Error(err))
	}
	res.Success(c, "情侣关系绑定成功")
}

// CoupleInvitation2
// @Tags Relationship
// @Summary 情侣关系绑定（通过交换注册码进行关系绑定）
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dto.RelationshipBinding true "未加密注册码，用户id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"情侣关系绑定成功"}"
// @Router /api/v1/relationship2 [post]
func (r Relationship) CoupleInvitation2(c *gin.Context) {
	var bindingParam dto.RelationshipBinding2
	if err := c.ShouldBindJSON(&bindingParam); err != nil {
		logger.Error("关系绑定请求参数有误", zap.Error(err))
		return
	}
	relationship := service.Relationship{}
	err := relationship.RelationBinding2(&bindingParam)
	if err != nil {
		logger.Error("情侣关系绑定失败", zap.Error(err))
	}
	res.Success(c, "情侣关系绑定成功")
}

// CoupleUnbound
// @Tags Relationship
// @Summary 情侣关系解绑
// @Security ApiKeyAuth
// @Produce  application/json
// @Param coupleId query int64 true "coupleID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"情侣关系解绑成功"}"
// @Router /api/v1/relationship [delete]
func (r *Relationship) CoupleUnbound(c *gin.Context) {
	coupleId := c.Query("coupleId")
	parseInt, err := strconv.ParseInt(coupleId, 10, 64)
	if err != nil {
		logger.Error("情侣关系解绑的请求参数有误", zap.Error(err))
		return
	}
	relationship := service.Relationship{}
	err = relationship.RelationUnbind(parseInt)
	if err != nil {
		logger.Error("情侣关系解绑失败", zap.Error(err))
	}
	res.Success(c, "情侣关系解绑成功")
}
