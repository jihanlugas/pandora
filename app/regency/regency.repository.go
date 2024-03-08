package regency

import (
	"fmt"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Regency, error)
	GetViewById(conn *gorm.DB, id string) (model.RegencyView, error)
	Page(conn *gorm.DB, req *request.PageRegency) ([]model.RegencyView, int64, error)
	List(conn *gorm.DB, req *request.ListRegency) ([]model.RegencyView, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Regency, error) {
	var err error
	var data model.Regency

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.RegencyView, error) {
	var err error
	var data model.RegencyView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Page(conn *gorm.DB, req *request.PageRegency) ([]model.RegencyView, int64, error) {
	var err error
	var data []model.RegencyView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(regency_name) LIKE LOWER(?)", "%"+req.RegencyName+"%")

	if req.ProvinceID != "" {
		query = query.Where("province_id = ?", req.ProvinceID)
	}

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "regency_name", "asc"))
	}
	err = query.Offset((req.GetPage() - 1) * req.GetLimit()).
		Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (r repository) List(conn *gorm.DB, req *request.ListRegency) ([]model.RegencyView, error) {
	var err error
	var data []model.RegencyView

	query := conn.Model(&data).
		Where("LOWER(regency_name) LIKE LOWER(?)", "%"+req.RegencyName+"%")

	if req.ProvinceID != "" {
		query = query.Where("province_id = ?", req.ProvinceID)
	}

	query = query.Order(fmt.Sprintf("%s %s", "regency_name", "asc"))

	err = query.Offset(req.GetLimit()).
		Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func NewRepository() Repository {
	return repository{}
}
