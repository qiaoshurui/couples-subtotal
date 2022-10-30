package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"time"
)

type User struct {
	ID                    int64     `json:"id"`
	UserName              string    `json:"userName"`
	Password              string    `json:"password"`
	NickName              string    `json:"nickName"`              //用户昵称
	Birthday              time.Time `json:"birthday"`              //出生日期
	Email                 string    `json:"email"`                 //用户邮箱
	Phone                 string    `json:"phone"`                 //用户手机号
	RegistrationCode      string    `json:"registrationCode"`      //注册码
	EncryptedRegistration string    `json:"encryptedRegistration"` //加密的注册码
	HeaderImg             string    `json:"headerImg"`             //用户头像
	CreatedAt             time.Time `json:"createdAt"`             //创建时间
	UpdatedAt             time.Time `json:"updatedAt"`             //更新时间
	IsDeleted             int8      `json:"isDeleted"`             //是否删除
}

type SimpleUser struct {
	ID        int64  `json:"id"`
	NickName  string `json:"nickName"`  //用户昵称
	HeaderImg string `json:"headerImg"` //用户头像
}

func (u *User) TableName() string {
	return "users"
}

func GetEmptyUser() *User {
	return new(User)
}

// SignRequest 注册请求参数
type SignRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	NickName string `json:"nickName"`
	Birthday int64  `json:"birthday"`
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
	err = global.Gorm.Where("user_name=? AND password=? AND is_deleted=0", user.UserName, user.Password).First(&user).Error
	return err
}

// ChangePassword 密码修改
func ChangePassword(mode *ChangePasswordRequest) (err error) {
	err = global.Gorm.Where("user_name=?", mode.UserName).Updates(&User{Password: encryptPassword(mode.NewPassword)}).Error
	return err
}

//func (u *User) GetUserById(id int64) (*User, error) {
//	var user *User
//	db := global.Gorm.Table(u.TableName()).First(&user, id)
//	return user, db.Error
//}

//func (u *User) GetUserById(id int64) (*User, error) {
//	db := global.Gorm.Table(u.TableName()).First(&u, id)
//	return u, db.Error
//}

func (u *User) GetUserById(id int64) error {
	db := global.Gorm.Table(u.TableName()).First(&u, id)
	return db.Error
}

func (u *User) GetSimpleUserById(id int64) (*SimpleUser, error) {
	var user *SimpleUser
	db := global.Gorm.Table(u.TableName()).Select("id,header_img,nick_name").First(&user, id)
	return user, db.Error
}
func (u *User) GetUserId(code string) error {
	db := global.Gorm.Table(u.TableName()).First(&u, "registration_code=?", code)
	return db.Error
}
