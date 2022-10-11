package service

import (
	model "github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model/request"
)

type User struct{}

func (u *User) SignUp(p *request.SignUp) (err error) {
	//判断用户是否存在
	if err = model.CheckUserExist(p.Username); err != nil {
		//查询数据库出错
		return err
	}
	//保存进数据库
	user := &model.User{
		UserName: p.Username,
		Password: p.Password,
		Email:    p.Email,
		Phone:    p.Phone,
	}
	err = model.InsertUser(user)
	return nil
}
func (u *User) Login(p *request.Login) (err error) {
	user := &model.User{
		Phone:    p.Phone,
		Password: p.Password,
	}
	//查询手机号和密码是否正确
	if err = model.LoginUser(user); err != nil {
		//手机号或密码错误
		return err
	}
	return
}
