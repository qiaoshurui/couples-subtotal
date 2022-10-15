package service

import (
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
)

// todo DynamicService
type Dynamic struct{}

func (d *Dynamic) AddDynamic(p *dto.AddDynamic) (err error) {
	dynamic := &model.Dynamic{
		Content: p.Content,
		UserId:  p.UserId,
		Status:  p.Status,
	}
	emptyDynamic := model.GetEmptyDynamic()
	err = emptyDynamic.InsertDynamic(dynamic)
	if err != nil {
		return errors.Wrapf(err, "新增动态失败 Content：%v", p.Content)
	}
	return nil
}
func (d *Dynamic) DeleteDynamic(data *model.Dynamic) (err error) {
	emptyDynamic := model.GetEmptyDynamic()
	err = emptyDynamic.DeleteDynamic(data)
	if err != nil {
		return errors.Wrap(err, "删除动态失败")
	}
	return nil
}
func (d *Dynamic) UpdateDynamic(data *model.Dynamic) (err error) {
	emptyDynamic := model.GetEmptyDynamic()
	err = emptyDynamic.UpdateDynamic(data)
	if err != nil {
		return errors.Wrap(err, "更新动态失败")
	}
	return nil
}

func (d *Dynamic) GetDynamicList(data *dto.GetDynamicList) (dynamicList []*dto.DynamicListInfo, err error) {
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
