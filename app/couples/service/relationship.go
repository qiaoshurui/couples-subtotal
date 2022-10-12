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
