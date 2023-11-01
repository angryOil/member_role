package service

import (
	"context"
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
