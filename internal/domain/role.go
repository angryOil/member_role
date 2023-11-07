package domain

import (
	"errors"
	"member_role/internal/domain/vo"
	"time"
)

var _ Role = (*role)(nil)

type Role interface {
	ValidCreate() error
	ValidUpdate() error

	Update(cafeRoleIds string) Role

	ToInfo() vo.Info
	ToDetail() vo.Detail
	ToUpdate() vo.Update
}

type role struct {
	id          int
	memberId    int
	cafeId      int
	cafeRoleIds string
	createdAt   time.Time
}

func (r *role) ToInfo() vo.Info {
	return vo.Info{
		Id:          r.id,
		CafeRoleIds: r.cafeRoleIds,
	}
}

func (r *role) ToDetail() vo.Detail {
	return vo.Detail{
		ID:           r.id,
		MemberId:     r.memberId,
		CafeRolesIds: r.cafeRoleIds,
	}
}

func (r *role) ToUpdate() vo.Update {
	return vo.Update{
		ID:           r.id,
		MemberId:     r.memberId,
		CafeId:       r.cafeId,
		CafeRolesIds: r.cafeRoleIds,
		CreatedAt:    r.createdAt,
	}
}

const (
	InvalidCafeId   = "invalid cafe id"
	InvalidId       = "invalid id"
	InvalidMemberId = "invalid member id"
)

func (r *role) ValidCreate() error {
	if r.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if r.memberId < 1 {
		return errors.New(InvalidMemberId)
	}
	return nil
}

func (r *role) ValidUpdate() error {
	if r.id < 1 {
		return errors.New(InvalidId)
	}
	if r.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if r.memberId < 1 {
		return errors.New(InvalidMemberId)
	}
	return nil
}

func (r *role) Update(cafeRoleIds string) Role {
	r.cafeRoleIds = cafeRoleIds
	return r
}
