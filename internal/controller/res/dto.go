package res

import "member_role/internal/domain"

type ListTotalDto[T any] struct {
	Contents []T `json:"contents"`
	Total    int `json:"total"`
}

func NewListTotalDto[T any](contents []T, total int) ListTotalDto[T] {
	return ListTotalDto[T]{
		Contents: contents,
		Total:    total,
	}
}

// 기본룰입니다 특정 사용자 권한을 확인할때 사용됩니다.

type MemberRole struct {
	Id         int    `json:"id,omitempty"`
	CafeRoleId string `json:"cafe_role_ids,omitempty"`
}

func ToDto(d domain.Role) MemberRole {
	return MemberRole{
		Id:         d.Id,
		CafeRoleId: d.CafeRoleIds,
	}
}

// 전체 사용자 권한을 확인할때 사용됩니다.

type MemberDetailRole struct {
	Id         int    `json:"id"`
	CafeRoleId string `json:"cafe_role_ids"`
	MemberId   int    `json:"member_id"`
}

func ToDetailList(domains []domain.Role) []MemberDetailRole {
	results := make([]MemberDetailRole, len(domains))
	for i, d := range domains {
		results[i] = MemberDetailRole{
			Id:         d.Id,
			MemberId:   d.MemberId,
			CafeRoleId: d.CafeRoleIds,
		}
	}
	return results
}
