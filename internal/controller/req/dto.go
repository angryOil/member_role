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

type PatchDto struct {
	CafeRoleIds string `json:"cafe_role_ids"`
}

func (d PatchDto) ToDomain(memberId, cafeId, id int) domain.Role {
	return domain.Role{
		Id:          id,
		MemberId:    memberId,
		CafeId:      cafeId,
		CafeRoleIds: d.CafeRoleIds,
	}
}
