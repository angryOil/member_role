package repository

import (
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"member_role/internal/domain"
	"member_role/internal/domain/vo"
	"member_role/internal/page"
	"member_role/internal/repository/model"
	"member_role/internal/repository/request"
)

type Repository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}

const (
	InternalServerError = "internal server error"
)

func (r Repository) CreateRole(ctx context.Context, cr request.CreateRole) error {
	cModel := model.ToCreateModel(cr)
	_, err := r.db.NewInsert().Model(&cModel).Exec(ctx)

	if err != nil {
		log.Println("CreateRole NewInsert err: ", err)
		return err
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
		log.Println("GetDetailListList NewSelect err: ", err)
		return []domain.Role{}, 0, errors.New("internal server error")
	}

	return model.ToDomainList(models), total, nil
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

func (r Repository) Save(ctx context.Context, id int, validFunc func(domains []domain.Role) (domain.Role, error), meredFunc func(d domain.Role) (vo.Update, error)) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Update beginTx err: ", err)
		return errors.New(InternalServerError)
	}

	var models []model.Role
	err = tx.NewSelect().Model(&models).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("Update NewSelect err: ", err)
		return errors.New(InternalServerError)
	}

	domains := model.ToDomainList(models)

	validDomain, err := validFunc(domains)
	if err != nil {
		return err
	}
	uVo, err := meredFunc(validDomain)
	if err != nil {
		return err
	}
	m := model.ToUpdateModel(request.Save{
		Id:          uVo.ID,
		MemberId:    uVo.MemberId,
		CafeId:      uVo.CafeId,
		CafeRoleIds: uVo.CafeRolesIds,
		CreatedAt:   uVo.CreatedAt,
	})

	_, err = tx.NewInsert().Model(&m).
		On("conflict (id) do update").Exec(ctx)
	if err != nil {
		log.Println("Save NewInsert err: ", err)
		return errors.New(InternalServerError)
	}
	err = tx.Commit()
	if err != nil {
		log.Println("Save Commit err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}
