package district

import (
	"fmt"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.District, error)
	GetViewById(conn *gorm.DB, id string) (model.DistrictView, error)
	Page(conn *gorm.DB, req *request.PageDistrict) ([]model.DistrictView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.District, error) {
	var err error
	var data model.District

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.DistrictView, error) {
	var err error
	var data model.DistrictView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Page(conn *gorm.DB, req *request.PageDistrict) ([]model.DistrictView, int64, error) {
	var err error
	var data []model.DistrictView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(district_name) LIKE LOWER(?)", "%"+req.DistrictName+"%")

	if req.ProvinceID != "" {
		query = query.Where("province_id = ?", req.ProvinceID)
	}

	if req.RegencyID != "" {
		query = query.Where("regency_id = ?", req.RegencyID)
	}

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "district_name", "asc"))
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
