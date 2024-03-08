package regency

import (
	"github.com/jihanlugas/pandora/db"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
)

type Usecase interface {
	GetById(id string) (model.RegencyView, error)
	Page(req *request.PageRegency) ([]model.RegencyView, int64, error)
	List(req *request.ListRegency) ([]model.RegencyView, error)
}

type usecaseRegency struct {
	repo Repository
}

func (u usecaseRegency) GetById(id string) (model.RegencyView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseRegency) Page(req *request.PageRegency) ([]model.RegencyView, int64, error) {
	var err error
	var data []model.RegencyView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (u usecaseRegency) List(req *request.ListRegency) ([]model.RegencyView, error) {
	var err error
	var data []model.RegencyView

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err = u.repo.List(conn, req)
	if err != nil {
		return data, err
	}

	return data, err
}

func NewUsecase(repo Repository) Usecase {
	return usecaseRegency{
		repo: repo,
	}
}
