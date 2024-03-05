package district

import (
	"github.com/jihanlugas/pandora/db"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
)

type Usecase interface {
	GetById(id string) (model.DistrictView, error)
	Page(req *request.PageDistrict) ([]model.DistrictView, int64, error)
}

type usecaseDistrict struct {
	repo Repository
}

func (u usecaseDistrict) GetById(id string) (model.DistrictView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseDistrict) Page(req *request.PageDistrict) ([]model.DistrictView, int64, error) {
	var err error
	var data []model.DistrictView
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
	return usecaseDistrict{
		repo: repo,
	}
}
