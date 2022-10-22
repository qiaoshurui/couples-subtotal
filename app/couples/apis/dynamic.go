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

// AddDynamic
// @Tags Dynamic
// @Summary 动态新增
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body dto.AddDynamic true "内容、用户id、动态状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"动态新增成功"}"
// @Router /api/v1/dynamic  [post]
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

// DeleteDynamic
// @Tags Dynamic
// @Summary 动态删除
// @Security ApiKeyAuth
// @Produce application/json
// @Param id query int64 true "ID"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"动态删除成功"}"
// @Router /api/v1/dynamic [delete]
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
// @Tags Dynamic
// @Summary 动态修改
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body dto.UpdateDynamic true "id,内容,状态,更新时间"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"动态修改成功"}"
// @Router /api/v1/dynamic [put]
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
// @Tags Dynamic
// @Summary 获取动态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.GetDynamicList true "页码, 每页大小,选填：动态内容"
// @Success 200 {object} dto.DynamicListInfo "分页获取动态列表,返回包括列表,总数"
// @Router /api/v1/dynamic/list [get]
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
