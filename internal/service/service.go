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

func (s Service) GetListByMemberId(ctx context.Context, cafeId int, memberId int) ([]domain.Role, error) {
	domains, err := s.repo.GetListByMemberId(ctx, cafeId, memberId)
	return domains, err
}

func (s Service) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]domain.Role, int, error) {
	domains, total, err := s.repo.GetList(ctx, cafeId, reqPage)
	return domains, total, err
}

func (s Service) Upsert(ctx context.Context, d domain.Role) error {
	if d.CafeRoleIds == "" {
		return errors.New("no cafe_role ids")
	}
	err := s.repo.Upsert(ctx, d)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, memberId int, id int) error {
	err := s.repo.Delete(ctx, cafeId, memberId, id)
	return err
}
