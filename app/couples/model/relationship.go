package model

import (
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"time"
)

type Relationship struct {
	ID           int       `json:"id"`
	CoupleId     int       `json:"couple_id"`     //情侣Aid
	PersonId     int       `json:"person_id"`     //情侣Bid
	MemorialDate int       `json:"memorial_date"` //纪念日
	CreatedAt    time.Time `json:"created_at"`    //创建时间
	UpdatedAt    time.Time `json:"updated_at"`    //更新时间
	IsDeleted    int8      `json:"is_deleted"`    //是否删除
}
type RelationshipRequest struct {
	ID int `json:"id"`
}
type Detail struct {
	ID        int    `json:"id"`
	NickName  string `json:"nick_name"`  //用户昵称
	HeaderImg string `json:"header_img"` //用户头像
}
type MiddleDisplayed struct {
	CoupleId     int `json:"couple_id"`     //情侣Aid
	PersonId     int `json:"person_id"`     //情侣Bid
	MemorialDate int `json:"memorial_date"` //纪念日
}

func GetDisplay(id int) (err error) {
	var ADetail Detail
	var middleDisplay MiddleDisplayed
	var BDetail Detail
	err = global.Gorm.Model(&User{}).Select("id,header_img,nick_name").Where("id=?", id).First(&ADetail).Error
	err = global.Gorm.Model(&Relationship{}).Select("couple_id,person_id,memorial_date").Where("couple_id=?", id).First(&middleDisplay).Error
	err = global.Gorm.Model(&User{}).Select("id,header_img,nick_name").Where("id=?", middleDisplay.PersonId).First(&BDetail).Error
	return err
}
