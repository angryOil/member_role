package req

import (
	"member_role/internal/domain"
	"time"
)

type CreateDto struct {
	CafeRoleId int `json:"cafe_role_id"`
}

func (d CreateDto) ToDomain(memberId, cafeId int) domain.Role {
	return domain.Role{
		MemberId:   memberId,
		CafeId:     cafeId,
		CafeRoleId: d.CafeRoleId,
		CreatedAt:  time.Now(),
	}
}
