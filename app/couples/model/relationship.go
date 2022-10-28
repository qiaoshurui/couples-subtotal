package model

import (
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Relationship struct {
	ID           int64                 `json:"id"`
	CoupleId     int64                 `json:"coupleId"`                                 //情侣Aid
	PersonId     int64                 `json:"personId"`                                 //情侣Bid
	MemorialDate time.Time             `json:"memorialDate"`                             //纪念日期
	CreatedAt    time.Time             `json:"createdAt"`                                //创建时间
	UpdatedAt    time.Time             `json:"updatedAt"`                                //更新时间
	DeletedAt    time.Time             `json:"deletedAt"`                                //删除时间
	IsDeleted    soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // 使用 `1` `0` 标识是否删除
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
func (r *Relationship) InsertRelationship(relationship *Relationship) error {
	db := global.Gorm.Create(&relationship)
	return db.Error
}
func (r *Relationship) Delete(relationship *Relationship) error {
	db := global.Gorm.Where("couple_id=? OR person_id=?", relationship.CoupleId, relationship.CoupleId).Delete(&relationship)
	return db.Error

}
