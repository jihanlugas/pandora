package village

import (
	"fmt"
	"github.com/jihanlugas/pandora/model"
	"github.com/jihanlugas/pandora/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Village, error)
	GetViewById(conn *gorm.DB, id string) (model.VillageView, error)
	Page(conn *gorm.DB, req *request.PageVillage) ([]model.VillageView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Village, error) {
	var err error
	var data model.Village

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.VillageView, error) {
	var err error
	var data model.VillageView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Page(conn *gorm.DB, req *request.PageVillage) ([]model.VillageView, int64, error) {
	var err error
	var data []model.VillageView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(village_name) LIKE LOWER(?)", "%"+req.VillageName+"%")

	if req.ProvinceID != "" {
		query = query.Where("province_id = ?", req.ProvinceID)
	}

	if req.RegencyID != "" {
		query = query.Where("regency_id = ?", req.RegencyID)
	}

	if req.DistrictID != "" {
		query = query.Where("district_id = ?", req.DistrictID)
	}

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "village_name", "asc"))
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
