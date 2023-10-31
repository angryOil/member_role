package domain

import "time"

type Role struct {
	Id         int
	MemberId   int
	CafeId     int
	CafeRoleId int
	CreatedAt  time.Time
}
