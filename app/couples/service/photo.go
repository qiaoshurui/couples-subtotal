package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"io"
	"time"
)

type Photo struct{}

func (p *Photo) UploadTencent(filePath string, file io.Reader) error {
	_, err := global.Client.Object.Put(context.Background(), filePath, file, nil)
	if err != nil {
		return errors.Wrapf(err, "照片上传至腾讯云失败 FilePath：%v", filePath)
	}
	return nil
}
func (p *Photo) AddPhotoAlbum(data *dto.AddPhotoAlbum) error {
	photoAlbum := &model.PhotoAlbum{
		Name:      data.Name,
		OwnerId:   _MyId,
		Type:      data.Type,
		CreatedAt: time.Now(),
	}
	emptyPhotoAlbum := model.GetEmptyPhotoAlbum()
	err := emptyPhotoAlbum.InsertPhotoAlbum(photoAlbum)
	if err != nil {
		return errors.Wrapf(err, "新增相册失败 Name：%v", data.Name)
	}
	return nil
}
