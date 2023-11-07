package model

import (
	"github.com/uptrace/bun"
	"member_role/internal/domain"
	"member_role/internal/repository/request"
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

func ToCreateModel(cm request.CreateRole) Role {
	return Role{
		MemberId:    cm.MemberId,
		CafeId:      cm.CafeId,
		CafeRoleIds: cm.CafeRoleIds,
		CreatedAt:   cm.CreatedAt,
	}
}

func ToUpdateModel(s request.Save) Role {
	return Role{
		Id:          s.Id,
		MemberId:    s.MemberId,
		CafeId:      s.CafeId,
		CafeRoleIds: s.CafeRoleIds,
		CreatedAt:   s.CreatedAt,
	}
}

func ToDomainList(models []Role) []domain.Role {
	results := make([]domain.Role, len(models))
	for i, m := range models {
		results[i] = domain.NewRoleBuilder().
			Id(m.Id).
			MemberId(m.MemberId).
			CafeId(m.CafeId).
			CafeRoleIds(m.CafeRoleIds).
			CreatedAt(m.CreatedAt).
			Build()
	}
	return results
}
