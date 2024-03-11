package ktp

import (
	"github.com/jihanlugas/pandora/app/jwt"
	"github.com/jihanlugas/pandora/db"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
)

type Usecase interface {
	GetById(id string) (model.KtpView, error)
	Create(loginUser jwt.UserLogin, req *request.CreateKtp) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdateKtp) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PageKtp) ([]model.KtpView, int64, error)
}

type usecaseKtp struct {
	repo Repository
}

func (u usecaseKtp) GetById(id string) (model.KtpView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseKtp) Create(loginUser jwt.UserLogin, req *request.CreateKtp) error {
	var err error
	var data model.Ktp

	data = model.Ktp{
		Nik:              req.Nik,
		Nama:             req.Nama,
		TempatLahir:      req.TempatLahir,
		TanggalLahir:     req.TanggalLahir,
		JenisKelamin:     req.JenisKelamin,
		ProvinceID:       req.ProvinceID,
		RegencyID:        req.RegencyID,
		DistrictID:       req.DistrictID,
		VillageID:        req.VillageID,
		Alamat:           req.Alamat,
		Rtrw:             req.Rtrw,
		Pekerjaan:        req.Pekerjaan,
		StatusPerkawinan: req.StatusPerkawinan,
		Kewarganegaraan:  req.Kewarganegaraan,
		BerlakuHingga:    req.BerlakuHingga,
		PhotoId:          "",
		CreateBy:         loginUser.UserID,
		UpdateBy:         loginUser.UserID,
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	err = u.repo.Create(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseKtp) Update(loginUser jwt.UserLogin, id string, req *request.UpdateKtp) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.Nik = req.Nik
	data.Nama = req.Nama
	data.TempatLahir = req.TempatLahir
	data.TanggalLahir = req.TanggalLahir
	data.JenisKelamin = req.JenisKelamin
	data.ProvinceID = req.ProvinceID
	data.RegencyID = req.RegencyID
	data.DistrictID = req.DistrictID
	data.VillageID = req.VillageID
	data.Alamat = req.Alamat
	data.Rtrw = req.Rtrw
	data.Pekerjaan = req.Pekerjaan
	data.StatusPerkawinan = req.StatusPerkawinan
	data.Kewarganegaraan = req.Kewarganegaraan
	data.BerlakuHingga = req.BerlakuHingga
	data.UpdateBy = loginUser.UserID

	tx := conn.Begin()

	err = u.repo.Update(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseKtp) Delete(loginUser jwt.UserLogin, id string) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	tx := conn.Begin()

	err = u.repo.Delete(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseKtp) Page(req *request.PageKtp) ([]model.KtpView, int64, error) {
	var err error
	var data []model.KtpView
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
	return usecaseKtp{
		repo: repo,
	}
}
