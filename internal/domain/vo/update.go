package vo

import "time"

type Update struct {
	ID           int       // 멤버권한 아이디 식별자
	MemberId     int       // 멤버 식별자
	CafeId       int       // 카페 식별자
	CafeRolesIds string    // 해당 멤버 카페 권한 아이디 리스트
	CreatedAt    time.Time // 생성일
}
