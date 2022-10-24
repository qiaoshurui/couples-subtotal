package service

import (
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
	"github.com/qiaoshurui/couples-subtotal/app/couples/service/dto"
	"github.com/qiaoshurui/couples-subtotal/common/utils"
	"time"
)

type Relationship struct{}

func (r *Relationship) GetRelationship(userId int64) (*dto.CouplesInfo, error) {
	user := model.GetEmptyUser()
	err := user.GetUserById(userId)

	relationship := model.GetEmptyRelationship()
	err = relationship.GetByCoupleAId(userId)

	lover := model.GetEmptyUser()
	err = lover.GetUserById(relationship.PersonId)

	if err != nil {
		return nil, errors.Wrapf(err, "查询失败 userId：%v", userId)
	}

	memorialDay := utils.SubDays(time.Now(), relationship.MemorialDate)

	couplesInfo := &dto.CouplesInfo{
		UserId:         user.ID,
		UserNickName:   user.NickName,
		UserHeaderImg:  user.HeaderImg,
		LoverId:        lover.ID,
		LoverNickName:  lover.NickName,
		LoverHeaderImg: lover.HeaderImg,
		MemorialDay:    memorialDay,
	}

	return couplesInfo, nil
}
func (r *Relationship) RelationBinding(data *dto.RelationshipBinding) (err error) {
	//注册码解码
	decryptionCode := utils.PasswordDecryption(data.RegistrationCode)
	//通过注册码找到邀请人id
	user := model.GetEmptyUser()
	err = user.GetUserId(decryptionCode)
	//添加数据到关系表
	relationship := &model.Relationship{
		CoupleId:  _MyId,
		PersonId:  user.ID,
		CreatedAt: time.Now(),
	}
	emptyRelationship := model.GetEmptyRelationship()
	err = emptyRelationship.InsertRelationship(relationship)
	if err != nil {
		return errors.Wrapf(err, "情侣关系绑定失败 coupleId：%v", data.UserId)
	}
	return nil
}
func (r *Relationship) RelationUnbind(coupleId int64) error {
	relationship := &model.Relationship{CoupleId: coupleId}
	emptyRelationship := model.GetEmptyRelationship()
	err := emptyRelationship.Delete(relationship)
	if err != nil {
		return errors.Wrapf(err, "情侣关系解绑失败 coupleId：%v", coupleId)
	}
	return nil
}
