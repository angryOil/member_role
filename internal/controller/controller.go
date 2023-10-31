package controller

import (
	"context"
	"member_role/internal/controller/req"
	"member_role/internal/controller/res"
	"member_role/internal/service"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Create(ctx context.Context, cafeId int, memberId int, cd req.CreateDto) error {
	d := cd.ToDomain(memberId, cafeId)
	err := c.s.Create(ctx, d)
	return err
}

func (c Controller) GetListByMemberId(ctx context.Context, cafeId int, memberId int) ([]res.MemberRole, error) {
	domains, err := c.s.GetListByMemberId(ctx, cafeId, memberId)
	if err != nil {
		return []res.MemberRole{}, err
	}
	return res.ToDtoList(domains), nil
}