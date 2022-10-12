package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/res"
)

type Relationship struct {
	api.Api
}

// Display 情侣关系展示页面
func (r Relationship) Display(c *gin.Context) {
	////获取参数校验
	//var s model.RelationshipRequest
	//if err := c.ShouldBindJSON(&s); err != nil {
	//	//logger.Error("注册请求参数有误")
	//	res.ParamError(c)
	//	return
	//}
	//Relationship := service.Relationship{}
	//if err := Relationship.GetRelationship(&s); err != nil {
	//	//logger.Error(err)
	//	//res.Error(err)
	//	//return
	//}
	//返回响应
	res.Success(c, "用户注册成功")
}
