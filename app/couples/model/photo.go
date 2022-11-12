package model

import (
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"gorm.io/plugin/soft_delete"
	"time"
)

type Photo struct {
	Id        int64                 `json:"id"`
	UserId    int64                 `json:"userId"`                                   //创建者id
	AlbumId   int64                 `json:"albumId"`                                  //相册id
	ImgUrl    string                `json:"imgUrl"`                                   //照片路径
	CreatedAt time.Time             `json:"createdAt"`                                //创建时间
	UpdatedAt time.Time             `json:"updatedAt"`                                //更新时间
	DeletedAt time.Time             `json:"deletedAt"`                                //删除时间
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt"` // 使用 `1` `0` 标识是否删除
}

func (p *Photo) TableName() string {
	return "photo"
}

func GetEmptyPhoto() *Photo {
	return new(Photo)
}

func (p *Photo) InsertPhoto(photo *Photo) error {
	db := global.Gorm.Select("user_id", "album_id", "img_url", "created_at").Create(&photo)
	return db.Error
}

func (p *Photo) GetPhotoList(data *dto.PhotoList) ([]*dto.PhotoListRes, error) {
	limit := data.PageSize
	offset := data.PageSize * (data.Page - 1)
	var photoList []*dto.PhotoListRes
	db := global.Gorm.Table(p.TableName()).Where("album_id = ?", data.AlbumId).Limit(limit).Offset(offset).Find(&photoList)
	return photoList, db.Error
}

func (p *Photo) DeletePhotos(ids []int64) error {
	db := global.Gorm.Delete(&p, ids)
	return db.Error
}

func (p *Photo) GetUrls(ids []int64) ([]string, error) {
	var urls []string
	db := global.Gorm.Table(p.TableName()).Select("img_url").Where("id IN ?", ids).Find(&urls)
	return urls, db.Error
}
func (p *Photo) DeletePhotoByAlbum(albumIds []int64) error {
	db := global.Gorm.Where("album_id IN ?", albumIds).Delete(&p)
	return db.Error
}

func (p *Photo) GetPhotoCount(id int64) (int64, error) {
	var count int64
	db := global.Gorm.Table(p.TableName()).Where("album_id=?", id).Count(&count)
	return count, db.Error
}
