package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qiaoshurui/couples-subtotal/app/couples/apis"
)

func RegisterUserRouter(v1 *gin.RouterGroup) {
	user := &apis.User{}
	{
		v1.POST("/sign", user.SignUpHandler)            //用户注册
		v1.POST("/login", user.LoginHandler)            //用户登录
		v1.POST("/changePassword", user.ChangePassword) //修改密码
	}
}
