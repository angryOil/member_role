package service

import (
	"context"
	"member_role/internal/domain"
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
