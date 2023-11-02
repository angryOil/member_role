package req

import (
	"member_role/internal/domain"
	"time"
)

type CreateDto struct {
	CafeRoleIds string `json:"cafe_role_ids"`
}

func (d CreateDto) ToDomain(memberId, cafeId int) domain.Role {
	return domain.Role{
		MemberId:    memberId,
		CafeId:      cafeId,
		CafeRoleIds: d.CafeRoleIds,
		CreatedAt:   time.Now(),
	}
}

type PutDto struct {
	CafeRoleIds string `json:"cafe_role_ids"`
}

func (d PutDto) ToDomain(memberId, cafeId int) domain.Role {
	return domain.Role{
		MemberId:    memberId,
		CafeId:      cafeId,
		CafeRoleIds: d.CafeRoleIds,
		CreatedAt:   time.Now(),
	}
}
