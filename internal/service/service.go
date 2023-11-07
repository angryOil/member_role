package service

import (
	"context"
	"errors"
	"member_role/internal/domain"
	"member_role/internal/domain/vo"
	"member_role/internal/page"
	"member_role/internal/repository"
	"member_role/internal/repository/request"
	"member_role/internal/service/req"
	"member_role/internal/service/res"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

const (
	NotFoundError = "not found"
)

func (s Service) CreateRole(ctx context.Context, cr req.CreateDto) error {
	createdAt := time.Now()
	err := domain.NewRoleBuilder().
		MemberId(cr.MemberId).
		CafeId(cr.CafeId).
		CafeRoleIds(cr.CafeRoleIds).
		CreatedAt(createdAt).
		Build().
		ValidCreate()
	if err != nil {
		return err
	}
	err = s.repo.CreateRole(ctx, request.CreateRole{
		MemberId:    cr.MemberId,
		CafeId:      cr.CafeId,
		CafeRoleIds: cr.CafeRoleIds,
		CreatedAt:   createdAt,
	})
	return err
}

func (s Service) GetListByMemberId(ctx context.Context, cafeId int, memberId int) (res.GetListByMemberId, error) {
	domains, err := s.repo.GetListByMemberId(ctx, cafeId, memberId)
	if err != nil || len(domains) == 0 {
		return res.GetListByMemberId{}, err
	}
	info := domains[0].ToInfo()
	return res.GetListByMemberId{
		Id:         info.Id,
		CafeRoleId: info.CafeRoleIds,
	}, err
}

func (s Service) GetDetailListList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.GetDetailList, int, error) {
	domains, total, err := s.repo.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.GetDetailList{}, 0, err
	}
	return convertDomainArrToGetListArr(domains), total, nil
}

func convertDomainArrToGetListArr(domains []domain.Role) []res.GetDetailList {
	result := make([]res.GetDetailList, len(domains))
	for i, d := range domains {
		vo := d.ToDetail()
		result[i] = res.GetDetailList{
			Id:          vo.ID,
			CafeRoleIds: vo.CafeRolesIds,
			MemberId:    vo.MemberId,
		}
	}
	return result
}

func (s Service) Save(ctx context.Context, u req.Save) error {
	id, memberId, cafeId := u.Id, u.MemberId, u.CafeId
	cafeRoleIds := u.CafeRoleIds

	err := domain.NewRoleBuilder().
		Id(id).
		MemberId(memberId).
		CafeId(cafeId).
		CafeRoleIds(cafeRoleIds).
		Build().
		ValidUpdate()
	if err != nil {
		return err
	}
	err = s.repo.Save(ctx, u.Id,
		func(domains []domain.Role) (domain.Role, error) {
			if len(domains) != 1 {
				return domain.NewRoleBuilder().Build(), errors.New(NotFoundError)
			}
			return domains[0], nil
		},
		func(d domain.Role) (vo.Update, error) {
			u := d.Update(cafeRoleIds)
			err := u.ValidUpdate()
			if err != nil {
				return vo.Update{}, err
			}
			return u.ToUpdate(), nil
		},
	)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, memberId int, id int) error {
	err := s.repo.Delete(ctx, cafeId, memberId, id)
	return err
}
