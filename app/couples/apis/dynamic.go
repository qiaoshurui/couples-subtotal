package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"go.uber.org/zap"
	"time"
)

type Dynamic struct {
	api.Api
}

// AddDynamic 动态新增
func (d *Dynamic) AddDynamic(c *gin.Context) {
	var s dto.AddDynamic
	if err := c.ShouldBindJSON(&s); err != nil {
		logger.Error("动态新增请求参数有误", zap.Error(err))
		return
	}
	dynamic := service.Dynamic{}
	err := dynamic.AddDynamic(&model.Dynamic{
		Content:   s.Content,
		UserId:    s.UserId,
		Status:    s.Status,
		CreatedAt: time.Now(),
	})
	if err != nil {
		logger.Error("动态新增失败", zap.Error(err))
	}
	res.OkWithMessage(c, "动态新增成功", nil)

}

// DeleteDynamic 动态删除
func (d *Dynamic) DeleteDynamic(c *gin.Context) {
	var s dto.DeleteDynamic
	if err := c.ShouldBindJSON(&s); err != nil {
		logger.Error("动态删除请求参数有误", zap.Error(err))
		return
	}
	dynamic := service.Dynamic{}
	err := dynamic.DeleteDynamic(&model.Dynamic{
		ID:        s.Id,
		IsDeleted: 1,
	})
	if err != nil {
		logger.Error("动态删除失败", zap.Error(err))
	}
	res.OkWithMessage(c, "动态删除成功", nil)
}

// UpdateDynamic 动态修改
func (d *Dynamic) UpdateDynamic(c *gin.Context) {
	var s dto.UpdateDynamic
	if err := c.ShouldBindJSON(&s); err != nil {
		logger.Error("动态修改请求参数有误", zap.Error(err))
	}
	dynamic := service.Dynamic{}
	err := dynamic.UpdateDynamic(&model.Dynamic{
		ID:        s.Id,
		Content:   s.Content,
		Status:    s.Status,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		logger.Error("动态修改失败", zap.Error(err))
	}
	res.OkWithMessage(c, "动态修改成功", nil)
}

// GetDynamicList 动态查找列表
func (d *Dynamic) GetDynamicList(c *gin.Context) {
	var s dto.GetDynamicList
	if err := c.ShouldBindJSON(&s); err != nil {
		logger.Error("动态查找请求参数有误", zap.Error(err))
	}
	dynamic := service.Dynamic{}
	dynamicList, err := dynamic.GetDynamicList(&s)
	if err != nil {
		logger.Error("动态查找失败", zap.Error(err))
	}
	res.OkWithMessage(c, "动态查找列表展示成功", dynamicList)
}
