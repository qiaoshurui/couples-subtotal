package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"github.com/tencentyun/cos-go-sdk-v5"
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
func (p *Photo) GetPhotoList(data *dto.PhotoList) ([]*dto.PhotoListRes, error) {
	emptyPhoto := model.GetEmptyPhoto()
	photoList, err := emptyPhoto.GetPhotoList(data)
	if err != nil {
		return nil, errors.Wrapf(err, "查询相册列表失败 albumId: %v", data.AlbumId)
	}
	return photoList, nil
}
func (p *Photo) DeletePhoto(ids []int64) error {
	emptyPhoto := model.GetEmptyPhoto()
	err := emptyPhoto.DeletePhotos(ids)
	if err != nil {
		return errors.Wrapf(err, "照片删除失败 albumIds: %v", ids)
	}
	return nil
}
func (p *Photo) DeleteCosRecord(ids []int64) error {
	emptyPhoto := model.GetEmptyPhoto()
	keys, err := emptyPhoto.GetUrls(ids)
	if err != nil {
		return errors.Wrapf(err, "查找照片的key失败 albumIds: %v", ids)
	}
	var obs []cos.Object
	for _, key := range keys {
		obs = append(obs, cos.Object{
			Key: key})
	}
	opt := &cos.ObjectDeleteMultiOptions{
		Objects: obs,
	}
	_, _, err = global.CosClient.Object.DeleteMulti(context.Background(), opt)
	if err != nil {
		return errors.Wrapf(err, "照片删除失败 albumIds: %v", ids)
	}
	return nil
}
