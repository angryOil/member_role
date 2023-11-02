package model

import (
	"github.com/uptrace/bun"
	"member_role/internal/domain"
	"time"
)

type Role struct {
	bun.BaseModel `bun:"table:member_role,alias:mr"`

	Id          int       `bun:"id,pk,autoincrement"`
	MemberId    int       `bun:"member_id,notnull"`
	CafeId      int       `bun:"cafe_id,notnull"`
	CafeRoleIds string    `bun:"cafe_role_ids,notnull"`
	CreatedAt   time.Time `bun:"created_at"`
}

func ToModel(d domain.Role) Role {
	return Role{
		Id:          d.Id,
		MemberId:    d.MemberId,
		CafeId:      d.CafeId,
		CafeRoleIds: d.CafeRoleIds,
		CreatedAt:   d.CreatedAt,
	}
}

func ToDomainList(models []Role) []domain.Role {
	results := make([]domain.Role, len(models))
	for i, m := range models {
		results[i] = domain.Role{
			Id:          m.Id,
			MemberId:    m.MemberId,
			CafeId:      m.CafeId,
			CafeRoleIds: m.CafeRoleIds,
			CreatedAt:   m.CreatedAt,
		}
	}
	return results
}
