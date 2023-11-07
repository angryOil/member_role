package request

import "time"

type CreateRole struct {
	MemberId    int
	CafeId      int
	CafeRoleIds string
	CreatedAt   time.Time
}

type Save struct {
	Id          int
	MemberId    int
	CafeId      int
	CafeRoleIds string
	CreatedAt   time.Time
}
