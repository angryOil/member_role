package domain

import "time"

var _ RoleBuilder = (*roleBuilder)(nil)

func NewRoleBuilder() RoleBuilder {
	return &roleBuilder{}
}

type RoleBuilder interface {
	Id(id int) RoleBuilder
	MemberId(memberId int) RoleBuilder
	CafeId(cafeId int) RoleBuilder
	CafeRoleIds(cafeRoleIds string) RoleBuilder
	CreatedAt(createdAt time.Time) RoleBuilder
	Build() Role
}

type roleBuilder struct {
	id          int
	memberId    int
	cafeId      int
	cafeRoleIds string
	createdAt   time.Time
}

func (r *roleBuilder) Build() Role {
	return &role{
		id:          r.id,
		memberId:    r.memberId,
		cafeId:      r.cafeId,
		cafeRoleIds: r.cafeRoleIds,
		createdAt:   r.createdAt,
	}
}

func (r *roleBuilder) Id(id int) RoleBuilder {
	r.id = id
	return r
}

func (r *roleBuilder) MemberId(memberId int) RoleBuilder {
	r.memberId = memberId
	return r
}

func (r *roleBuilder) CafeId(cafeId int) RoleBuilder {
	r.cafeId = cafeId
	return r
}

func (r *roleBuilder) CafeRoleIds(cafeRoleIds string) RoleBuilder {
	r.cafeRoleIds = cafeRoleIds
	return r
}

func (r *roleBuilder) CreatedAt(createdAt time.Time) RoleBuilder {
	r.createdAt = createdAt
	return r
}
