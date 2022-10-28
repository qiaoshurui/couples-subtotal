package model

import (
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
