package res

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

// 전체 사용자 권한을 확인할때 사용됩니다.

type MemberDetailRole struct {
	Id          int    `json:"id"`
	CafeRoleIds string `json:"cafe_role_ids"`
	MemberId    int    `json:"member_id"`
}
