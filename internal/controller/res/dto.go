package res

import "member_role/internal/domain"

// 기본룰입니다 특정 사용자 권한을 확인할때 사용됩니다.

type MemberRole struct {
	Id         int `json:"id"`
	CafeRoleId int `json:"cafe_role_id"`
}

// 전체 사용자 권한을 확인할때 사용됩니다.

type MemberDetailRole struct {
	Id         int `json:"id"`
	CafeRoleId int `json:"cafe_role_id"`
	MemberId   int `json:"member_id"`
}

func ToDtoList(domains []domain.Role) []MemberRole {
	results := make([]MemberRole, len(domains))
	for i, d := range domains {
		results[i] = MemberRole{
			Id:         d.Id,
			CafeRoleId: d.CafeRoleId,
		}
	}
	return results
}
