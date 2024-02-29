package ktp

import (
	"fmt"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Ktp, error)
	GetViewById(conn *gorm.DB, id string) (model.KtpView, error)
	Create(conn *gorm.DB, data model.Ktp) error
	Update(conn *gorm.DB, data model.Ktp) error
	Delete(conn *gorm.DB, data model.Ktp) error
	Page(conn *gorm.DB, req *request.PageKtp) ([]model.KtpView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Ktp, error) {
	var err error
	var data model.Ktp

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.KtpView, error) {
	var err error
	var data model.KtpView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Ktp) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Ktp) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Ktp) error {
	return conn.Delete(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PageKtp) ([]model.KtpView, int64, error) {
	var err error
	var data []model.KtpView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(nik) LIKE LOWER(?)", "%"+req.Nik+"%").
		Where("LOWER(nama) LIKE LOWER(?)", "%"+req.Nama+"%").
		Where("LOWER(alamat) LIKE LOWER(?)", "%"+req.Alamat+"%").
		Where("LOWER(kelurahan_desa) LIKE LOWER(?)", "%"+req.KelurahanDesa+"%").
		Where("LOWER(Kecamatan) LIKE LOWER(?)", "%"+req.Kecamatan+"%").
		Where("LOWER(kabupaten_kota) LIKE LOWER(?)", "%"+req.KabupatenKota+"%").
		Where("LOWER(provinsi) LIKE LOWER(?)", "%"+req.Provinsi+"%").
		Where("LOWER(status_perkawinan) LIKE LOWER(?)", "%"+req.StatusPerkawinan+"%").
		Where("LOWER(kewarganegaraan) LIKE LOWER(?)", "%"+req.Kewarganegaraan+"%")

	if req.CreateBy != "" {
		query = query.Where("create_by = ?", req.CreateBy)
	}

	if req.JenisKelamin != "" {
		query = query.Where("jenis_kelamin = ?", req.JenisKelamin)
	}

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "nama", "asc"))
	}
	err = query.Offset((req.GetPage() - 1) * req.GetLimit()).
		Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewRepository() Repository {
	return repository{}
}
