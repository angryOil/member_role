package repository

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"member_role/internal/domain"
	"member_role/internal/page"
	"member_role/internal/repository/model"
	"strings"
)

type Repository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}

func (r Repository) Create(ctx context.Context, d domain.Role) error {
	m := model.ToModel(d)
	_, err := r.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return err
		}
		log.Println("Create NewInsert err: ", err)
		return errors.New("internal server error")
	}
	return nil
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

func (r Repository) Patch(ctx context.Context, cafeId int, memberId int, id int,
	validFunc func(domains []domain.Role) (domain.Role, error),
	mergeFunc func(d domain.Role) domain.Role) error {
	var models []model.Role
	err := r.db.NewSelect().Model(&models).Where("cafe_id = ? and member_id = ? and id = ?", cafeId, memberId, id).
		Scan(ctx)
	if err != nil {
		log.Println("Patch NewSelect err: ", err)
		return errors.New("internal server error")
	}

	validDomain, err := validFunc(model.ToDomainList(models))
	if err != nil {
		return err
	}

	mergedDomain := mergeFunc(validDomain)

	m := model.ToModel(mergedDomain)
	_, err = r.db.NewInsert().Model(&m).On("CONFLICT (id) DO UPDATE").Exec(ctx)
	if err != nil {
		log.Println("Patch NewInsert err: ", err)
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
