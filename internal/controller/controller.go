package controller

import (
	"context"
	"member_role/internal/controller/req"
	"member_role/internal/controller/res"
	"member_role/internal/page"
	"member_role/internal/service"
	req2 "member_role/internal/service/req"
	res2 "member_role/internal/service/res"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}

func (c Controller) CreateRole(ctx context.Context, memberId int, cafeId int, cR req.CreateDto) error {
	err := c.s.CreateRole(ctx, req2.CreateDto{
		MemberId:    memberId,
		CafeId:      cafeId,
		CafeRoleIds: cR.CafeRoleIds,
	})
	return err
}

func (c Controller) GetListByMemberId(ctx context.Context, cafeId int, memberId int) (res.MemberRole, error) {
	dto, err := c.s.GetListByMemberId(ctx, cafeId, memberId)
	if err != nil {
		return res.MemberRole{}, err
	}
	return res.MemberRole{
		Id:         dto.Id,
		CafeRoleId: dto.CafeRoleId,
	}, nil
}

func (c Controller) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.MemberDetailRole, int, error) {
	infoArr, total, err := c.s.GetDetailListList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.MemberDetailRole{}, 0, err
	}
	return convertListInfoArrToDetailRoleArr(infoArr), total, nil
}

func convertListInfoArrToDetailRoleArr(listArr []res2.GetDetailList) []res.MemberDetailRole {
	result := make([]res.MemberDetailRole, len(listArr))
	for i, l := range listArr {
		result[i] = res.MemberDetailRole{
			Id:          l.Id,
			CafeRoleIds: l.CafeRoleIds,
			MemberId:    l.MemberId,
		}
	}
	return result
}

func (c Controller) Upsert(ctx context.Context, id int, cafeId int, memberId int, pDto req.PutDto) error {
	err := c.s.Save(ctx, req2.Save{
		Id:          id,
		MemberId:    memberId,
		CafeId:      cafeId,
		CafeRoleIds: pDto.CafeRoleIds,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, memberId int, id int) error {
	err := c.s.Delete(ctx, cafeId, memberId, id)
	return err
}
