package dto

import "time"

type CouplesInfo struct {
	UserId        int64  `json:"userId"`
	UserNickName  string `json:"userNickName"`
	UserHeaderImg string `json:"userHeaderImg"`

	LoverId        int64  `json:"loverId"`
	LoverNickName  string `json:"loverNickName"`
	LoverHeaderImg string `json:"loverHeaderImg"`

	MemorialDay int `json:"memorialDay"` //纪念日已过去天数
}
type RelationshipBinding struct {
	RegistrationCodeEncrypt string    `json:"registrationCodeEncrypt"` //加密后的注册码
	MemorialDate            time.Time `json:"memorialDate"`            //纪念日
	UserId                  int64     `json:"userId"`
}
type RelationshipBinding2 struct {
	RegistrationCodeDecryption string    `json:"registrationCodeDecryption"` //解密后的注册码
	MemorialDate               time.Time `json:"memorialDate"`               //纪念日
	UserId                     int64     `json:"userId"`
}
