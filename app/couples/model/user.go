package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	NickName  string    `json:"nick_name"`  //用户昵称
	Birthday  time.Time `json:"birthday"`   //出生日期
	Email     string    `json:"email"`      //用户邮箱
	Phone     string    `json:"phone"`      //用户手机号
	HeaderImg string    `json:"header_img"` //用户头像
	CreatedAt time.Time `json:"created_at"` //创建时间
	UpdatedAt time.Time `json:"updated_at"` //更新时间
	IsDeleted int8      `json:"is_deleted"` //是否删除
}

// SignRequest 注册请求参数
type SignRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// LoginRequest 登录请求参数
type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// ChangePasswordRequest 用户密码修改
type ChangePasswordRequest struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) (err error) {
	var count int64
	err = global.Gorm.Model(&User{}).Where("user_name=? AND is_deleted=0", username).Count(&count).Error
	if count > 0 {
		return errors.New("用户名已注册")
	}
	return nil
}

// InsertUser 用户注册
func InsertUser(user *User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//插入数据到数据库
	return global.Gorm.Create(&user).Error
}

// 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// LoginUser 用户登录
func LoginUser(user *User) (err error) {
	user.Password = encryptPassword(user.Password) //加密后的密码
	err = global.Gorm.Where("phone=? AND password=? AND is_deleted=0", user.Phone, user.Password).First(&user).Error
	return err
}

// ChangePassword 密码修改
func ChangePassword(mode *ChangePasswordRequest) (err error) {
	err = global.Gorm.Updates(&User{Password: encryptPassword(mode.NewPassword)}).Error
	return err
}
