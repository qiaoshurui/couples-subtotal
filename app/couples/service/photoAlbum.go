package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/global"
	"github.com/tencentyun/cos-go-sdk-v5"
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
func (p *PhotoAlbum) DeleteAlbum(ids []int64) error {
	//删除相册
	emptyPhotoAlbum := model.GetEmptyPhotoAlbum()
	err := emptyPhotoAlbum.DeleteAlbums(ids)
	if err != nil {
		return errors.Wrapf(err, "相册删除失败 albumIds: %v", ids)
	}
	//删除相册中的照片
	emptyPhoto := model.GetEmptyPhoto()
	err = emptyPhoto.DeletePhotoByAlbum(ids)
	if err != nil {
		return errors.Wrapf(err, "照片删除失败 albumIds: %v", ids)
	}
	return nil
}
func (p *PhotoAlbum) DeleteCosRecord(ids []int64) error {
	emptyPhotoAlbum := model.GetEmptyPhotoAlbum()
	keys, err := emptyPhotoAlbum.GetUrls(ids)
	if err != nil {
		return errors.Wrapf(err, "查找相册的key失败 albumIds: %v", ids)
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
		return errors.Wrapf(err, "相册删除失败 albumIds: %v", ids)
	}
	return nil
}
func (p *PhotoAlbum) GetAlbumList(data *dto.AlbumListReq) (albumListRes []*dto.AlbumListRes, err error) {
	photoAlbum := model.GetEmptyPhotoAlbum()
	albumLists, err := photoAlbum.GetAlbumList(data, _MyId)
	if err != nil {
		return nil, errors.Wrapf(err, "查询相册列表失败 Type: %v", data.Type)
	}
	for _, albumList := range albumLists {
		emptyPhoto := model.GetEmptyPhoto()
		photoCount, err := emptyPhoto.GetPhotoCount(albumList.Id)
		if err != nil {
			return nil, errors.Wrapf(err, "查询该相册的照片个数失败 AlbumId：%v", albumList.Id)
		}
		albumListInfo := &dto.AlbumListRes{
			Id:         albumList.Id,
			Name:       albumList.Name,
			Type:       albumList.Type,
			AlbumUrl:   albumList.AlbumUrl,
			PhotoCount: photoCount,
		}
		albumListRes = append(albumListRes, albumListInfo)
	}
	return
}
