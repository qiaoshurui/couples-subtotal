package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"strconv"
	"strings"
	"time"
)

type PhotoAlbum struct{}

func (p *PhotoAlbum) AddPhotoAlbum(data *dto.AddPhotoAlbum, key string, uuid string) error {
	photoAlbum := &model.PhotoAlbum{
		Name:      data.Name,
		OwnerId:   _MyId,
		Type:      data.Type,
		AlbumUrl:  key,
		Uuid:      uuid,
		CreatedAt: time.Now(),
	}
	emptyPhotoAlbum := model.GetEmptyPhotoAlbum()
	err := emptyPhotoAlbum.InsertPhotoAlbum(photoAlbum)
	if err != nil {
		return errors.Wrapf(err, "新增相册失败 Name：%v", data.Name)
	}
	return nil
}
func (p *PhotoAlbum) UploadTencent(num int8, uuid string) (string, error) {
	var key string
	if num == 1 {
		key = "user/" + strconv.Itoa(_MyId) + "/" + uuid + "/"
	} else {
		emptyRelationship := model.GetEmptyRelationship()
		err := emptyRelationship.GetCoupleId(4)
		fmt.Println(emptyRelationship.ID)
		if err != nil {
			return "", errors.Wrapf(err, "查询该用户的情侣id失败 UserId：%v", _MyId)
		}
		key = "couple/" + strconv.FormatInt(emptyRelationship.ID, 10) + "/" + uuid + "/"
	}
	_, err := global.CosClient.Object.Put(context.Background(), key, strings.NewReader(""), nil)
	if err != nil {
		return "", errors.Wrapf(err, "创建相册到腾讯云失败 Key:%v", key)
	}
	return key, nil
}
