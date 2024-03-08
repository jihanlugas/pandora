package province

import (
	"github.com/jihanlugas/pandora/db"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
)

type Usecase interface {
	GetById(id string) (model.ProvinceView, error)
	Page(req *request.PageProvince) ([]model.ProvinceView, int64, error)
	List(req *request.ListProvince) ([]model.ProvinceView, error)
}

type usecaseProvince struct {
	repo Repository
}

func (u usecaseProvince) GetById(id string) (model.ProvinceView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseProvince) Page(req *request.PageProvince) ([]model.ProvinceView, int64, error) {
	var err error
	var data []model.ProvinceView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (u usecaseProvince) List(req *request.ListProvince) ([]model.ProvinceView, error) {
	var err error
	var data []model.ProvinceView

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err = u.repo.List(conn, req)
	if err != nil {
		return data, err
	}

	return data, err
}

func NewUsecase(repo Repository) Usecase {
	return usecaseProvince{
		repo: repo,
	}
}
