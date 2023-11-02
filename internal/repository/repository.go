package repository

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"member_role/internal/domain"
	"member_role/internal/page"
	"member_role/internal/repository/model"
)

type Repository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}

func (r Repository) GetListByMemberId(ctx context.Context, cafeId int, memberId int) ([]domain.Role, error) {
	var models []model.Role
	err := r.db.NewSelect().Model(&models).Where("cafe_id = ? and member_id = ?", cafeId, memberId).Scan(ctx)
	if err != nil {
		log.Println("GetListByMemberId NewSelect err: ", err)
		return []domain.Role{}, errors.New("internal server error")
	}
	return model.ToDomainList(models), nil
}

func (r Repository) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]domain.Role, int, error) {
	var models []model.Role
	total, err := r.db.NewSelect().Model(&models).Where("cafe_id = ?", cafeId).Limit(reqPage.Size).Offset(reqPage.OffSet).Order("id desc").ScanAndCount(ctx)

	if err != nil {
		log.Println("GetList NewSelect err: ", err)
		return []domain.Role{}, 0, errors.New("internal server error")
	}

	return model.ToDomainList(models), total, nil
}

func (r Repository) Upsert(ctx context.Context, d domain.Role) error {
	m := model.ToModel(d)

	_, err := r.db.NewInsert().Model(&m).
		On("CONFLICT (cafe_id,member_id) DO UPDATE").
		Set("cafe_role_ids = EXCLUDED.cafe_role_ids").
		Exec(ctx)

	if err != nil {
		log.Println("Upsert NewInsert err: ", err)
		return errors.New("internal server error")
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, cafeId int, memberId int, id int) error {
	var m model.Role
	_, err := r.db.NewDelete().Model(&m).Where("cafe_id = ? and member_id = ? and id = ?", cafeId, memberId, id).Exec(ctx)
	if err != nil {
		log.Println("Delete NewDelete err: ", err)
		return errors.New("internal server error")
	}
	return nil
}
