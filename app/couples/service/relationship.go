package service

import (
	"github.com/pkg/errors"
	"github.com/qiaoshurui/couples-subtotal/app/couples/model"
)

type Relationship struct{}

func (r *Relationship) GetRelationship(p *model.RelationshipRequest) (err error) {
	if err = model.GetDisplay(p.ID); err != nil {
		return errors.Wrap(err, "查询情侣关系内容失败")
	}
	return nil
}
