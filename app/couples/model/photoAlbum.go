package model

import (
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"gorm.io/plugin/soft_delete"
	"time"
)

type PhotoAlbum struct {
	Id        int64                 `json:"id"`
	Name      string                `json:"name"`                                     //相册名称
	OwnerId   int64                 `json:"ownerId"`                                  //创建人id
	Type      string                `json:"type"`                                     //相册可见类型(0 情侣相册；1 个人相册)
	CreatedAt time.Time             `json:"createdAt"`                                //创建时间
	UpdatedAt time.Time             `json:"updatedAt"`                                //更新时间
	DeletedAt time.Time             `json:"deletedAt"`                                //删除时间
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // 使用 `1` `0` 标识是否删除
}

func (p *PhotoAlbum) TableName() string {
	return "photo_album"
}
func GetEmptyPhotoAlbum() *PhotoAlbum {
	return new(PhotoAlbum)
}
func (p *PhotoAlbum) InsertPhotoAlbum(photoAlbum *PhotoAlbum) (err error) {
	return global.Gorm.Select("Name", "OwnerId", "Type", "CreatedAt").Create(photoAlbum).Error
}
