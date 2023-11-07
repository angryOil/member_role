package domain

import (
	"errors"
	"time"
)

var _ Role = (*role)(nil)

type Role interface {
	ValidCreate() error
	Update(cafeRoleIds string) Role
}

type role struct {
	id          int
	memberId    int
	cafeId      int
	cafeRoleIds string
	createdAt   time.Time
}

const (
	InvalidCafeId   = "invalid cafe id"
	InvalidMemberId = "invalid member id"
)

func (r *role) ValidCreate() error {
	if r.CafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if r.MemberId < 1 {
		return errors.New(InvalidMemberId)
	}
	return nil
}

func (r *role) Update(cafeRoleIds string) Role {
	r.CafeRoleIds = cafeRoleIds
	return r
}
