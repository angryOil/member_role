package controller

import (
	"context"
	"member_role/internal/controller/req"
	"member_role/internal/controller/res"
	"member_role/internal/page"
	"member_role/internal/service"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetListByMemberId(ctx context.Context, cafeId int, memberId int) (res.MemberRole, error) {
	domains, err := c.s.GetListByMemberId(ctx, cafeId, memberId)
	if err != nil {
		return res.MemberRole{}, err
	}
	return res.ToDto(domains), nil
}

func (c Controller) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.MemberDetailRole, int, error) {
	domains, total, err := c.s.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.MemberDetailRole{}, 0, err
	}
	return res.ToDetailList(domains), total, nil
}

func (c Controller) Upsert(ctx context.Context, cafeId int, memberId int, pDto req.PutDto) error {
	d := pDto.ToDomain(memberId, cafeId)
	err := c.s.Upsert(ctx, d)
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, memberId int, id int) error {
	err := c.s.Delete(ctx, cafeId, memberId, id)
	return err
}
