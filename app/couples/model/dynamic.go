package model

import (
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"time"
)

type Dynamic struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`    //动态内容
	UserId    int64     `json:"user_id"`    //动态发布者id
	Status    int8      `json:"status"`     //动态状态（0 双方可见;1 仅自己可见）
	CreatedAt time.Time `json:"created_at"` //创建时间
	UpdatedAt time.Time `json:"updated_at"` //更新时间
	IsDeleted int8      `json:"is_deleted"` //是否删除
}

func (d *Dynamic) TableName() string {
	return "dynamic"
}
func GetEmptyDynamic() *Dynamic {
	return new(Dynamic)
}
func (d *Dynamic) InsertDynamic(dynamic *Dynamic) (err error) {
	//插入数据到数据库
	return global.Gorm.Create(&dynamic).Error
}

// todo 思考指针用法
func (d *Dynamic) Insert() (err error) {
	//插入数据到数据库
	return global.Gorm.Create(&d).Error
}

// todo 命名冗余系列   删除一般根据id删除   了解gorm逻辑删除写法
func (d *Dynamic) DeleteDynamic(dynamic *Dynamic) (err error) {
	return global.Gorm.Updates(&dynamic).Error
}
func (d *Dynamic) UpdateDynamic(dynamic *Dynamic) (err error) {
	return global.Gorm.Updates(&dynamic).Error
}
func (d *Dynamic) GetDynamicList(data *dto.GetDynamicList) ([]*Dynamic, error) {
	limit := data.PageSize
	offset := data.PageSize * (data.PageSize - 1)
	var dynamicList []*Dynamic
	db := global.Gorm.Table(d.TableName()).Where("content like %?% AND is_deleted=0", data.Content).Limit(limit).Offset(offset).Find(&dynamicList)
	return dynamicList, db.Error
}
