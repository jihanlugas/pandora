package village

import (
	"github.com/jihanlugas/pandora/db"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
)

type Usecase interface {
	GetById(id string) (model.VillageView, error)
	Page(req *request.PageVillage) ([]model.VillageView, int64, error)
}

type usecaseVillage struct {
	repo Repository
}

func (u usecaseVillage) GetById(id string) (model.VillageView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseVillage) Page(req *request.PageVillage) ([]model.VillageView, int64, error) {
	var err error
	var data []model.VillageView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewUsecase(repo Repository) Usecase {
	return usecaseVillage{
		repo: repo,
	}
}
