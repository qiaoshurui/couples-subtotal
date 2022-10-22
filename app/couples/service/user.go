package service

import (
	"github.com/pkg/errors"
	model "github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/common/utils"
)

type User struct{}

func (u *User) SignUp(p *model.SignRequest) (err error) {
	//判断用户是否存在
	if err = model.CheckUserExist(p.UserName); err != nil {
		//查询数据库出错
		return errors.Wrapf(err, "用户名已存在username:%s", p.UserName)
	}
	//保存进数据库
	registrationCode := utils.RegCodeCreat()
	encryptedRegistration := utils.PasswordEncryption(registrationCode)

	user := &model.User{
		UserName:              p.Phone,
		Password:              p.Password,
		Email:                 p.Email,
		Phone:                 p.Phone,
		RegistrationCode:      registrationCode,
		EncryptedRegistration: encryptedRegistration,
	}
	if err = model.InsertUser(user); err != nil {
		return errors.Wrap(err, "注册用户落库失败")
	}
	return nil
}
func (u *User) Login(p *model.LoginRequest) (err error) {
	user := &model.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	//查询用户名和密码是否正确
	if err = model.LoginUser(user); err != nil {
		//手机号或密码错误
		return errors.Wrapf(err, "手机号或用户密码错误 username: %s", p.UserName)
	}
	return nil
}
func (u *User) ChangePassword(p *model.ChangePasswordRequest) (err error) {
	user := &model.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	//查询用户名和原密码是否正确
	if err = model.LoginUser(user); err != nil {
		//手机号或密码错误
		return errors.Wrapf(err, "手机号或用户密码错误 username: %s", p.UserName)
	}
	//修改密码
	if err = model.ChangePassword(p); err != nil {
		return errors.Wrap(err, "修改用户密码失败")
	}
	return nil
}
