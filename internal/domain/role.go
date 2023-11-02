package domain

import "time"

type Role struct {
	Id          int
	MemberId    int
	CafeId      int
	CafeRoleIds string
	CreatedAt   time.Time
}
