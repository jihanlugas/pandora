package province

import (
	"fmt"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Province, error)
	GetViewById(conn *gorm.DB, id string) (model.ProvinceView, error)
	Page(conn *gorm.DB, req *request.PageProvince) ([]model.ProvinceView, int64, error)
	List(conn *gorm.DB, req *request.ListProvince) ([]model.ProvinceView, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Province, error) {
	var err error
	var data model.Province

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.ProvinceView, error) {
	var err error
	var data model.ProvinceView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Page(conn *gorm.DB, req *request.PageProvince) ([]model.ProvinceView, int64, error) {
	var err error
	var data []model.ProvinceView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(province_name) LIKE LOWER(?)", "%"+req.ProvinceName+"%")

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "province_name", "asc"))
	}
	err = query.Offset((req.GetPage() - 1) * req.GetLimit()).
		Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (r repository) List(conn *gorm.DB, req *request.ListProvince) ([]model.ProvinceView, error) {
	var err error
	var data []model.ProvinceView

	query := conn.Model(&data).
		Where("LOWER(province_name) LIKE LOWER(?)", "%"+req.ProvinceName+"%").
		Order(fmt.Sprintf("%s %s", "province_name", "asc"))

	err = query.Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func NewRepository() Repository {
	return repository{}
}
