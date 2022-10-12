package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service"
	"github.com/qiaoshurui/couples-subtotal/common/api"
	"github.com/qiaoshurui/couples-subtotal/common/logger"
	"github.com/qiaoshurui/couples-subtotal/common/res"
	"go.uber.org/zap"
)

type User struct {
	api.Api
}

// SignUpHandler 注册请求函数
func (u *User) SignUpHandler(c *gin.Context) {
	//获取参数校验
	var s model.SignRequest
	if err := c.ShouldBindJSON(&s); err != nil {
		//logger.Error("注册请求参数有误")
		res.ParamError(c)
		return
	}
	//业务处理
	singUser := service.User{}
	if err := singUser.SignUp(&s); err != nil {
		logger.Error("用户注册失败", zap.Error(err))
		res.Error(c, "用户注册失败", "用户注册失败")
		return
	}
	//返回响应
	res.Success(c, "用户注册成功")
}

// LoginHandler 登录请求函数
func (u *User) LoginHandler(c *gin.Context) {
	//获取参数校验
	var s model.LoginRequest
	if err := c.ShouldBindJSON(&s); err != nil {
		logger.Error("登录请求参数有误")
		res.ParamError(c)
		return
	}
	//业务处理
	singUser := service.User{}
	if err := singUser.Login(&s); err != nil {
		logger.Error("登陆失败! 用户名不存在或者密码错误!")
		return
	}
	//TODO:tokenManager
	//返回响应
	res.Success(c, "登录成功")
}

// ChangePassword 用户密码修改
func (u *User) ChangePassword(c *gin.Context) {
	var s model.ChangePasswordRequest
	//参数校验
	if err := c.ShouldBindJSON(&s); err != nil {
		logger.Error("密码修改请求参数有误")
		res.ParamError(c)
		return
	}
	//业务处理
	singUser := service.User{}
	if err := singUser.ChangePassword(&s); err != nil {
		logger.Error("修改失败,原密码与当前账户不符")
		return
	}
	res.Success(c, "用户密码修改成功")

}