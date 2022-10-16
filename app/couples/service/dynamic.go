package service

import (
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"time"
)

type DynamicService struct{}

const _MyId = 1

func (d *DynamicService) AddDynamic(data *dto.AddDynamic) (err error) {
	dynamic := &model.Dynamic{
		Content:   data.Content,
		UserId:    _MyId,
		Status:    data.Status,
		CreatedAt: time.Now(),
	}
	emptyDynamic := model.GetEmptyDynamic()
	err = emptyDynamic.InsertDynamic(dynamic)
	if err != nil {
		return errors.Wrapf(err, "新增动态失败 Content：%v", data.Content)
	}
	return nil
}

func (d *DynamicService) DeleteDynamic(id int64) (err error) {
	dynamic := &model.Dynamic{ID: id}
	emptyDynamic := model.GetEmptyDynamic()
	err = emptyDynamic.Delete(dynamic)
	if err != nil {
		return errors.Wrapf(err, "删除动态失败 ID:%v", id)
	}
	return nil
}

func (d *DynamicService) UpdateDynamic(data *dto.UpdateDynamic) (err error) {
	dynamic := &model.Dynamic{
		ID:        data.Id,
		Content:   data.Content,
		Status:    data.Status,
		UpdatedAt: time.Now(),
	}
	emptyDynamic := model.GetEmptyDynamic()
	err = emptyDynamic.UpdateDynamic(dynamic)
	if err != nil {
		return errors.Wrapf(err, "更新动态失败 ID:%v", data.Id)
	}
	return nil
}

func (d *DynamicService) GetDynamicList(data *dto.GetDynamicList) (dynamicList []*dto.DynamicListInfo, err error) {
	emptyDynamic := model.GetEmptyDynamic()
	dynamics, err := emptyDynamic.GetDynamicList(data)
	if err != nil {
		return nil, errors.Wrap(err, "查询动态列表失败")
	}
	for _, dynamic := range dynamics {
		user := model.GetEmptyUser()
		err = user.GetUserById(dynamic.UserId)
		dynamicDetail := &dto.DynamicListInfo{
			Id:           dynamic.ID,
			Content:      dynamic.Content,
			UserId:       dynamic.UserId,
			UserNickName: user.NickName,
		}
		dynamicList = append(dynamicList, dynamicDetail)
	}
	return
}
func (d *DynamicService) GetDynamicDetail(id int64) (*dto.SimpleDynamicDetail, error) {
	emptyDynamic := model.GetEmptyDynamic()
	dynamicDetail, err := emptyDynamic.GetDynamicDetail(id)
	if err != nil {
		return nil, errors.Wrapf(err, "查询动态详情失败 ID: %v", id)
	}
	return dynamicDetail, nil
}
