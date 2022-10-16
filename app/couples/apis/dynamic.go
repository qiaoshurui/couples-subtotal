package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"go.uber.org/zap"
	"strconv"
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
	dynamic := service.DynamicService{}
	err := dynamic.AddDynamic(&s)
	if err != nil {
		logger.Error("动态新增失败", zap.Error(err))
	}
	res.OkWithMessage(c, "动态新增成功", nil)

}

// DeleteDynamic 动态删除
func (d *Dynamic) DeleteDynamic(c *gin.Context) {
	dynamicId := c.Query("id")
	id, err := strconv.ParseInt(dynamicId, 10, 64)
	if err != nil {
		logger.Error("动态删除请求参数有误", zap.Error(err))
		return
	}
	dynamic := service.DynamicService{}
	err = dynamic.DeleteDynamic(id)
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
	dynamic := service.DynamicService{}
	err := dynamic.UpdateDynamic(&s)
	if err != nil {
		logger.Error("动态修改失败", zap.Error(err))
	}
	res.OkWithMessage(c, "动态修改成功", nil)
}

// GetDynamicList 查看动态列表
func (d *Dynamic) GetDynamicList(c *gin.Context) {
	var s dto.GetDynamicList
	if err := c.ShouldBindJSON(&s); err != nil {
		logger.Error("动态查找列表请求参数有误", zap.Error(err))
	}
	dynamic := service.DynamicService{}
	dynamicList, err := dynamic.GetDynamicList(&s)
	if err != nil {
		logger.Error("动态查找失败", zap.Error(err))
	}
	res.OkWithMessage(c, "动态查找列表展示成功", dynamicList)
}

// GetDynamicDetail 查看动态详情
func (d *Dynamic) GetDynamicDetail(c *gin.Context) {
	dynamicId := c.Param("id")
	parseInt, err := strconv.ParseInt(dynamicId, 10, 64)
	if err != nil {
		logger.Error("查看动态详情请求参数有误", zap.Error(err))
	}
	dynamicService := service.DynamicService{}
	getDynamic, err := dynamicService.GetDynamicDetail(parseInt)
	if err != nil {
		logger.Error("动态详情展示失败", zap.Error(err))
	}
	res.OkWithMessage(c, "动态详情展示成功", getDynamic)

}
