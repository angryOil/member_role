package service

import (
	"context"
	"errors"
	"member_role/internal/domain"
	"member_role/internal/page"
	"member_role/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(ctx context.Context, d domain.Role) error {
	err := s.repo.Create(ctx, d)
	return err
}

func (s Service) GetListByMemberId(ctx context.Context, cafeId int, memberId int) ([]domain.Role, error) {
	domains, err := s.repo.GetListByMemberId(ctx, cafeId, memberId)
	return domains, err
}

func (s Service) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]domain.Role, int, error) {
	domains, total, err := s.repo.GetList(ctx, cafeId, reqPage)
	return domains, total, err
}

func (s Service) Patch(ctx context.Context, reqD domain.Role) error {
	err := s.repo.Patch(ctx, reqD.CafeId, reqD.MemberId, reqD.Id,
		func(domains []domain.Role) (domain.Role, error) {
			if len(domains) == 0 {
				return domain.Role{}, errors.New("no rows")
			}
			return domains[0], nil
		},
		func(d domain.Role) domain.Role {
			return domain.Role{
				Id:          d.Id,
				MemberId:    d.MemberId,
				CafeId:      d.CafeId,
				CafeRoleIds: reqD.CafeRoleIds,
				CreatedAt:   d.CreatedAt,
			}
		},
	)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, memberId int, id int) error {
	err := s.repo.Delete(ctx, cafeId, memberId, id)
	return err
}
