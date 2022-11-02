package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"io"
	"time"
)

type Photo struct{}

func (p *Photo) UploadTencent(imgName string, albumId int64, file io.Reader) (string, error) {
	PhotoAlbum := model.GetEmptyPhotoAlbum()
	err := PhotoAlbum.GetAlbumUrl(albumId)
	if err != nil {
		return "", errors.Wrapf(err, "数据库查询相册路径失败 AlbumUrl：%v", PhotoAlbum.AlbumUrl)
	}
	key := PhotoAlbum.AlbumUrl + imgName
	_, err = global.CosClient.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		return "", errors.Wrapf(err, "照片上传至腾讯云失败 Key：%v", key)
	}
	return key, nil
}
func (p *Photo) AddPhoto(imgUrl string, albumId int64) error {
	photo := &model.Photo{
		UserId:    _MyId,
		AlbumId:   albumId,
		ImgUrl:    imgUrl,
		CreatedAt: time.Now(),
	}
	emptyPhoto := model.GetEmptyPhoto()
	err := emptyPhoto.InsertPhoto(photo)
	if err != nil {
		return errors.Wrapf(err, "照片上传至数据库失败 ImgURl：%v", imgUrl)
	}
	return nil
}
