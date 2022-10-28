package model

import (
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Dynamic struct {
	ID        int64                 `json:"id"`
	Content   string                `json:"content"`                                  //动态内容
	UserId    int64                 `json:"userId"`                                   //动态发布者id
	Status    int8                  `json:"status"`                                   //动态状态（0 双方可见;1 仅自己可见）
	CreatedAt time.Time             `json:"createdAt"`                                //创建时间
	UpdatedAt time.Time             `json:"updatedAt"`                                //更新时间
	DeletedAt time.Time             `json:"deletedAt"`                                //删除时间
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // 使用 `1` `0` 标识是否删除
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

func (d *Dynamic) Insert() (err error) {
	//插入数据到数据库
	return global.Gorm.Create(&d).Error
}

func (d *Dynamic) Delete(dynamic *Dynamic) (err error) {
	return global.Gorm.Delete(&dynamic).Error
}
func (d *Dynamic) UpdateDynamic(dynamic *Dynamic) (err error) {
	return global.Gorm.Updates(&dynamic).Error
}
func (d *Dynamic) GetDynamicList(data *dto.GetDynamicList) ([]*Dynamic, error) {
	limit := data.PageSize
	offset := data.PageSize * (data.Page - 1)
	var dynamicList []*Dynamic
	db := global.Gorm.Table(d.TableName()).Where("content like ? ", "%"+data.Content+"%").Limit(limit).Offset(offset).Find(&dynamicList)
	return dynamicList, db.Error
}

//func (d *Dynamic) GetDynamicList2(page, size int, content string) ([]*Dynamic, error) {
//	limit := size
//	offset := size * (page - 1)
//	var dynamicList []*Dynamic
//	db := global.Gorm.Table(d.TableName()).Where("content like ? ", "%"+content+"%").Limit(limit).Offset(offset).Find(&dynamicList)
//	return dynamicList, db.Error
//}

func (d *Dynamic) GetDynamicDetail(id int64) (*dto.SimpleDynamicDetail, error) {
	var dynamicDetail *dto.SimpleDynamicDetail
	db := global.Gorm.Table(d.TableName()).Select("id,content,user_id,created_at").First(&dynamicDetail, id)
	return dynamicDetail, db.Error
}
