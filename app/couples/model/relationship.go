package model

import (
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"time"
)

type Relationship struct {
	ID           int64     `json:"id"`
	CoupleId     int64     `json:"couple_id"`     //情侣Aid
	PersonId     int64     `json:"person_id"`     //情侣Bid
	MemorialDate time.Time `json:"memorial_date"` //纪念日
	CreatedAt    time.Time `json:"created_at"`    //创建时间
	UpdatedAt    time.Time `json:"updated_at"`    //更新时间
	IsDeleted    int8      `json:"is_deleted"`    //是否删除
}

func (r *Relationship) TableName() string {
	return "lovers_relationship"
}

func GetEmptyRelationship() *Relationship {
	return new(Relationship)
}

func (r *Relationship) GetByCoupleAId(coupleAId int64) error {
	db := global.Gorm.Model(r).Where("couple_id=?", coupleAId).First(&r)
	return db.Error
}
